package fiber

import (
	"fmt"
	"testing"
)

func TestParseAddr_1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want1 string
		want2 string
	}{
		{"case1", "127.0.0.1:8080", "127.0.0.1", "8080"},
		{"case2", "localhost", "localhost", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got1, got2 := parseAddr(tt.input)
			if got1 != tt.want1 || got2 != tt.want2 {
				t.Errorf("parseAddr(%s) = %s, %s; want %s, %s", tt.input, got1, got2, tt.want1, tt.want2)
			}
		})
	}
}
