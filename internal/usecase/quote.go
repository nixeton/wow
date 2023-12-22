package usecase

type QuoteRepository interface {
	GetRandomQuote() string
}

type QuoteUsecase struct {
	quoteRepository QuoteRepository
}

func NewQuoteUsecase(repository QuoteRepository) *QuoteUsecase {
	return &QuoteUsecase{
		quoteRepository: repository,
	}
}

func (uc *QuoteUsecase) GetQuote() string {
	return uc.quoteRepository.GetRandomQuote()
}
