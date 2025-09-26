package middleware

import (
	"fmt"
	"testing"
)

func TestValidateCSRFToken_1(t *testing.T) {
	type args struct {
		token       string
		clientToken string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				token:       "123456",
				clientToken: "123456",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				token:       "123456",
				clientToken: "654321",
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				token:       "",
				clientToken: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := validateCSRFToken(tt.args.token, tt.args.clientToken); got != tt.want {
				t.Errorf("validateCSRFToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
