package usecase

import (
	"testing"
	"wow/internal/data/repository"
)

func TestQuoteUsecase_GetQuote(t *testing.T) {
	type fields struct {
		quoteRepository QuoteRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Test GetQuote - success case",
			fields: fields{
				quoteRepository: repository.NewRepository(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &QuoteUsecase{
				quoteRepository: tt.fields.quoteRepository,
			}
			got, err := uc.GetQuote()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetQuote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 {
				t.Errorf("GetQuote() result must be greate than 0")
			}
		})
	}
}
