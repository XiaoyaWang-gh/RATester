package fiber

import (
	"fmt"
	"testing"
)

func TestparseAddr_1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want1 string
		want2 string
	}{
		{"no colon", "localhost", "localhost", ""},
		{"one colon", "localhost:8080", "localhost", "8080"},
		{"multiple colons", "localhost:8080:8081", "localhost:8080", "8081"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got1, got2 := parseAddr(tt.input)
			if got1 != tt.want1 {
				t.Errorf("parseAddr() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("parseAddr() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
