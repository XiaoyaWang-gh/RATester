package utils

import (
	"fmt"
	"net/smtp"
	"testing"
)

func TestAttachFile_1(t *testing.T) {
	e := &Email{
		Auth:     smtp.PlainAuth("", "user", "password", "host"),
		Identity: "identity",
		Username: "username",
		Password: "password",
		Host:     "host",
		Port:     587,
		From:     "from",
	}

	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "valid file",
			args:    []string{"testdata/test.txt"},
			wantErr: false,
		},
		{
			name:    "invalid file",
			args:    []string{"testdata/invalid.txt"},
			wantErr: true,
		},
		{
			name:    "empty file",
			args:    []string{""},
			wantErr: true,
		},
		{
			name:    "too many arguments",
			args:    []string{"testdata/test.txt", "id", "extra"},
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

			_, err := e.AttachFile(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("AttachFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
