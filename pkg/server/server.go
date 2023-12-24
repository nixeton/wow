package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"net"
	"strconv"
	"time"
	"wow/config"
	"wow/helpers"
	"wow/pkg/logger"
)

type ProceedPow interface {
	GetChallenge() []byte
	Verify(challenge, solution []byte) error
}

type ProceedQuote interface {
	GetQuote() (string, error)
}

type Server struct {
	config       *config.Config
	logger       *logger.Logger
	listener     net.Listener
	cancel       context.CancelFunc
	proceedPow   ProceedPow
	proceedQuote ProceedQuote
}

// NewServer creates a new server
func NewServer(
	config *config.Config,
	proceedPow ProceedPow,
	proceedQuote ProceedQuote,
	logger *logger.Logger,
) *Server {

	return &Server{
		config:       config,
		proceedPow:   proceedPow,
		proceedQuote: proceedQuote,
		logger:       logger,
	}
}

// Start starts the server
func (s *Server) Start(ctx context.Context) (err error) {
	ctx, s.cancel = context.WithCancel(ctx)
	defer s.cancel()

	lc := net.ListenConfig{
		KeepAlive: s.config.KeepAlive,
	}
	s.listener, err = lc.Listen(ctx, "tcp", s.config.Address)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	s.listen(ctx)

	return nil
}

// Stop stops the server
func (s *Server) Stop() {
	s.cancel()
}

// listen listens for incoming connections
func (s *Server) listen(ctx context.Context) {

	go func() {
		<-ctx.Done()
		err := s.listener.Close()

		if err != nil && !errors.Is(err, net.ErrClosed) {
			s.logger.Error("failed to close listener: ", err.Error())
		}
	}()

	for {
		conn, err := s.listener.Accept()
		if errors.Is(err, net.ErrClosed) {
			s.logger.Debug("listener closed")
			return
		} else if err != nil {
			s.logger.Error("failed to accept connection: ", err.Error())
			continue
		}

		if err = conn.SetDeadline(time.Now().Add(s.config.TCP.Deadline)); err != nil {
			s.logger.Error("failed to set deadline: ", err.Error())
			err := conn.Close()
			if err != nil {
				s.logger.Error("failed to close connection: ", err.Error())
			}
			continue
		}

		if err != nil {
			s.logger.Error("failed to set deadline: ", err.Error())
			err := conn.Close()
			if err != nil {
				s.logger.Error("failed to close connection: ", err.Error())
			}
			continue
		}

		go s.handle(conn)
	}
}

func (s *Server) handle(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			s.logger.Error("failed to close connection: ", err.Error())
		}
	}(conn)

	_, err := helpers.Read(conn)
	if err != nil {
		s.logger.Error("failed to read message: ", err.Error())
		return
	}

	// send difficulty
	diff := s.config.Difficulty
	diffs := []byte(strconv.FormatInt(int64(diff), 10))
	if err := helpers.Write(conn, diffs); err != nil {
		s.logger.Error("failed to write difficulty: ", err.Error())
		return
	}

	// send challenge
	challenge := s.proceedPow.GetChallenge()
	if err := helpers.Write(conn, challenge); err != nil {
		s.logger.Error("failed to write challenge: ", err.Error())
		return
	}
	log.Info().Msg("challenge: " + string(challenge))

	// receive nonce
	nonce, err := helpers.Read(conn)
	if err != nil {
		s.logger.Error("failed to read solution: ", err.Error())
		return
	}

	// verify nonce
	if err = s.proceedPow.Verify(challenge, nonce); err != nil {
		s.logger.Error("failed to verify solution: ", err.Error())
		return
	}

	// send quote
	quote, err := s.proceedQuote.GetQuote()
	if err != nil {
		s.logger.Error("failed to get quote")
		return
	}

	if err = helpers.Write(conn, []byte(quote)); err != nil {
		s.logger.Error("failed to write quote: ", err.Error())
		return
	}
}
