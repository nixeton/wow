package service

import "wow/internal/usecase"

type QuoteService struct {
	quoteUsecase *usecase.QuoteUsecase
}

func NewQuoteService(quoteUsecase *usecase.QuoteUsecase) *QuoteService {
	return &QuoteService{
		quoteUsecase: quoteUsecase,
	}
}

func (s *QuoteService) GetQuote() string {
	return s.quoteUsecase.GetQuote()
}
