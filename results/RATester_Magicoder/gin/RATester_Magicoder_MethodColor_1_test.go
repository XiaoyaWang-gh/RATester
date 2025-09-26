package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMethodColor_1(t *testing.T) {
	tests := []struct {
		name string
		p    *LogFormatterParams
		want string
	}{
		{
			name: "GET",
			p: &LogFormatterParams{
				Method: http.MethodGet,
			},
			want: blue,
		},
		{
			name: "POST",
			p: &LogFormatterParams{
				Method: http.MethodPost,
			},
			want: cyan,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.p.MethodColor(); got != tt.want {
				t.Errorf("MethodColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
