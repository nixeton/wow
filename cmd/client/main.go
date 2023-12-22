package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
	"strconv"
	"strings"
	"wow/helpers"
)

var difficulty int

const maxCount = 100000000000

func main() {
	client := NewClient()
	err := client.Run(context.Background(), maxCount)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = client.GetQuote(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}
}

type client struct {
}

func NewClient() *client {

	return &client{}
}

func (c *client) Run(ctx context.Context, count int) error {
	for i := 0; i < count; i++ {
		if ctx.Err() != nil {
			break
		}

		q, err := c.GetQuote(ctx)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(q))
		}
	}

	return nil
}

func (c *client) GetQuote(ctx context.Context) ([]byte, error) {
	var dialer net.Dialer
	conn, err := dialer.DialContext(ctx, "tcp", "localhost:8888")
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	if err := helpers.Write(conn, []byte("challenge")); err != nil {
		return nil, fmt.Errorf("send challenge request err: %w", err)
	}

	diffBytes, err := helpers.Read(conn)
	if err != nil {
		return nil, fmt.Errorf("receive difficulty err: %w", err)
	}
	difficulty, err = strconv.Atoi(string(diffBytes))

	challenge, err := helpers.Read(conn)
	if err != nil {
		return nil, fmt.Errorf("receive challenge err: %w", err)
	}

	_, n := solveChallenge(string(challenge))
	ns := strconv.FormatInt(n, 10)
	if err := helpers.Write(conn, []byte(ns)); err != nil {
		return nil, fmt.Errorf("send solution err: %w", err)
	}

	quote, err := helpers.Read(conn)
	if err != nil {
		return nil, fmt.Errorf("receive quote err: %w", err)
	}

	return quote, nil
}

func solveChallenge(challenge string) (string, int64) {
	var nonce int64
	var hash string
	prefix := strings.Repeat("0", difficulty)

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
