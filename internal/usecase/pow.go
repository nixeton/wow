package usecase

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	difficulty = 5
	prefix     = "0"
)

type PowUsecase struct {
}

func NewPowUsecase() *PowUsecase {
	return &PowUsecase{}
}

func (uc *PowUsecase) Verify(challenge string, nonce int) error {
	if !verify(challenge, nonce) {
		return errors.New("invalid nonce")
	}

	return nil
}

func (uc *PowUsecase) GetChallenge() []byte {
	token := generateChallenge()

	return []byte(token)
}

func generateChallenge() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	challenge := strconv.Itoa(rand.Int())
	return challenge
}

func verify(challenge string, nonce int) bool {
	prefix := strings.Repeat(prefix, difficulty)
	data := challenge + strconv.Itoa(nonce)

	hashBytes := sha256.Sum256([]byte(data))
	hash := hex.EncodeToString(hashBytes[:])

	return strings.HasPrefix(hash, prefix)
}
