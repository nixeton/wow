package usecase

import (
	"testing"
	"wow/internal/data/repository"
)

func TestQuoteUsecase_GetQuote(t *testing.T) {

	type fields struct {
	}

	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "GetRandomQuote",
			fields: fields{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository.NewRepository()

			uc := &QuoteUsecase{
				quoteRepository: r,
			}

			if got := uc.GetQuote(); len(got) == 0 {
				t.Errorf("GetQuote() = %v, want len > 0", got)
			}
		})
	}
}
