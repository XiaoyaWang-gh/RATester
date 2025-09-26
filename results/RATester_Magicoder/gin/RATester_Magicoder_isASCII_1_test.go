package gin

import (
	"fmt"
	"testing"
)

func TestisASCII_1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"ASCII", "Hello, World!", true},
		{"Non-ASCII", "Héllo, W�rld!", false},
		{"Empty", "", true},
		{"Only ASCII", "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", true},
		{"Only Non-ASCII", "éèêë", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isASCII(tt.s); got != tt.want {
				t.Errorf("isASCII() = %v, want %v", got, tt.want)
			}
		})
	}
}
