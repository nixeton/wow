package usecase

import (
	"testing"
	"wow/config"
)

func TestPowUsecase_GetChallenge(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test GetChallenge",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &PowUsecase{}
			got := uc.GetChallenge()
			if len(got) == 0 {
				t.Errorf("GetChallenge() = %v", got)
			}

			t.Log("GetChallenge() = ", uc.GetChallenge())
		})
	}
}

func TestPowUsecase_Verify(t *testing.T) {
	type args struct {
		challenge string
		nonce     int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Test Verify",
			args:    args{challenge: "1530108093636870175", nonce: 21387},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &PowUsecase{
				config: &config.Config{
					TCP: config.TCP{},
					Log: config.Log{},
					Pow: config.Pow{
						Difficulty: 5,
					},
				},
			}
			if err := uc.Verify(tt.args.challenge, tt.args.nonce); (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
