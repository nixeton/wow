package service

import (
	"strconv"
	"wow/internal/usecase"
)

type PowService struct {
	powUsecase *usecase.PowUsecase
}

func NewPowService(powUsecase *usecase.PowUsecase) *PowService {
	return &PowService{
		powUsecase: powUsecase,
	}
}

func (s *PowService) Verify(challenge []byte, nonce []byte) error {
	nonceInt, err := strconv.Atoi(string(nonce))
	if err != nil {
		return err
	}

	return s.powUsecase.Verify(string(challenge), nonceInt)
}

func (s *PowService) GetChallenge() []byte {
	return s.powUsecase.GetChallenge()
}
