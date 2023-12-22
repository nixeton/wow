package usecase

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"wow/config"
)

const (
	prefix = "0"
)

type PowUsecase struct {
	config *config.Config
}

func NewPowUsecase(config *config.Config) *PowUsecase {
	return &PowUsecase{
		config: config,
	}
}

// Verify verifies the nonce
func (uc *PowUsecase) Verify(challenge string, nonce int) error {
	if !verify(challenge, nonce, uc.config.Difficulty) {
		return errors.New("invalid nonce")
	}

	return nil
}

// GetChallenge returns a challenge
func (uc *PowUsecase) GetChallenge() []byte {
	token := generateChallenge()

	return []byte(token)
}

// GetQuote returns a quote
func generateChallenge() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	challenge := strconv.Itoa(rand.Int())
	return challenge
}

// verify verifies the nonce
func verify(challenge string, nonce int, difficulty int) bool {
	prefix := strings.Repeat(prefix, difficulty)
	data := challenge + strconv.Itoa(nonce)

	hashBytes := sha256.Sum256([]byte(data))
	hash := hex.EncodeToString(hashBytes[:])

	return strings.HasPrefix(hash, prefix)
}
