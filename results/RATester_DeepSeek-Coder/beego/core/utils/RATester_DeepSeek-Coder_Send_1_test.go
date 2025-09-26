package utils

import (
	"fmt"
	"testing"
)

func TestSend_1(t *testing.T) {
	type args struct {
		e *Email
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestSend_Success",
			args: args{
				e: &Email{
					Username: "test",
					Password: "test",
					Host:     "smtp.example.com",
					Port:     587,
					To:       []string{"test@example.com"},
				},
			},
			wantErr: false,
		},
		{
			name: "TestSend_NoToAddress",
			args: args{
				e: &Email{
					Username: "test",
					Password: "test",
					Host:     "smtp.example.com",
					Port:     587,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := tt.args.e.Send(); (err != nil) != tt.wantErr {
				t.Errorf("Email.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
