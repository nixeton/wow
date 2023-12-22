package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/rs/zerolog/log"
	"net"
	"strconv"
	"strings"
	"wow/helpers"
)

var difficulty int

// maxCount - max count of iterations
const maxCount = 1000
const prefix = "0"
const address = "localhost:8888"

func main() {
	client := NewClient()
	err := client.Run(context.Background(), maxCount)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	_, err = client.GetQuote(context.Background())
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
}

// solveChallenge - solve challenge
func solveChallenge(challenge string) (string, int64) {
	var nonce int64
	var hash string
	prefix := strings.Repeat(prefix, difficulty)

	for {
		nonce++
		data := challenge + strconv.FormatInt(nonce, 10)
		hashBytes := sha256.Sum256([]byte(data))
		hash = hex.EncodeToString(hashBytes[:])
		if strings.HasPrefix(hash, prefix) {
			break
		}
	}
	return hash, nonce
}

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

// Run - run client
func (c *Client) Run(ctx context.Context, count int) error {
	for i := 0; i < count; i++ {
		if ctx.Err() != nil {
			break
		}

		q, err := c.GetQuote(ctx)
		if err != nil {
			log.Error().Msg(err.Error())
		} else {
			fmt.Println(string(q))
		}
	}

	return nil
}

// GetQuote - get quote from server
func (c *Client) GetQuote(ctx context.Context) ([]byte, error) {
	var dialer net.Dialer
	conn, err := dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		return nil, fmt.Errorf("field to connect: %w", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Error().Msg(err.Error())
		}
	}()

	if err := helpers.Write(conn, []byte("challenge")); err != nil {
		return nil, fmt.Errorf("send challenge request err: %w", err)
	}

	// receive difficulty
	diffBytes, err := helpers.Read(conn)
	if err != nil {
		return nil, fmt.Errorf("receive difficulty err: %w", err)
	}
	difficulty, err = strconv.Atoi(string(diffBytes))

	// receive challenge
	challenge, err := helpers.Read(conn)
	if err != nil {
		return nil, fmt.Errorf("receive challenge err: %w", err)
	}

	// solve challenge
	_, n := solveChallenge(string(challenge))
	ns := strconv.FormatInt(n, 10)
	if err := helpers.Write(conn, []byte(ns)); err != nil {
		return nil, fmt.Errorf("send solution err: %w", err)
	}

	// receive quote
	quote, err := helpers.Read(conn)
	if err != nil {
		return nil, fmt.Errorf("receive quote err: %w", err)
	}

	return quote, nil
}
