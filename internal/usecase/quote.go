package usecase

type QuoteRepository interface {
	GetRandomQuote() (string, error)
}

type QuoteUsecase struct {
	quoteRepository QuoteRepository
}

func NewQuoteUsecase(repository QuoteRepository) *QuoteUsecase {
	return &QuoteUsecase{
		quoteRepository: repository,
	}
}

func (uc *QuoteUsecase) GetQuote() (string, error) {
	return uc.quoteRepository.GetRandomQuote()
}
