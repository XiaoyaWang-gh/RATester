package auth

import (
	"fmt"
	"testing"
)

func TestBasicUserParser_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantUser string
		wantPass string
		wantErr  bool
	}{
		{
			name:     "valid user",
			input:    "user:pass",
			wantUser: "user",
			wantPass: "pass",
			wantErr:  false,
		},
		{
			name:     "invalid user",
			input:    "user",
			wantUser: "",
			wantPass: "",
			wantErr:  true,
		},
		{
			name:     "empty user",
			input:    "",
			wantUser: "",
			wantPass: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotUser, gotPass, err := basicUserParser(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("basicUserParser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUser != tt.wantUser {
				t.Errorf("basicUserParser() gotUser = %v, want %v", gotUser, tt.wantUser)
			}
			if gotPass != tt.wantPass {
				t.Errorf("basicUserParser() gotPass = %v, want %v", gotPass, tt.wantPass)
			}
		})
	}
}
