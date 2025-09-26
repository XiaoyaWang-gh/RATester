package nathole

import (
	"fmt"
	"testing"
)

func TestNewTransactionID_1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test case 1",
			want: "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
		{
			name: "Test case 2",
			want: "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
		{
			name: "Test case 3",
			want: "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewTransactionID(); got != tt.want {
				t.Errorf("NewTransactionID() = %v, want %v", got, tt.want)
			}
		})
	}
}
