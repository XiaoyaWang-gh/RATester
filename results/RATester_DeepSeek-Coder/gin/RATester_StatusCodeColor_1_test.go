package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestStatusCodeColor_1(t *testing.T) {
	tests := []struct {
		name string
		p    *LogFormatterParams
		want string
	}{
		{
			name: "Test 1",
			p: &LogFormatterParams{
				StatusCode: http.StatusContinue,
			},
			want: white,
		},
		{
			name: "Test 2",
			p: &LogFormatterParams{
				StatusCode: http.StatusOK,
			},
			want: green,
		},
		{
			name: "Test 3",
			p: &LogFormatterParams{
				StatusCode: http.StatusMultipleChoices,
			},
			want: white,
		},
		{
			name: "Test 4",
			p: &LogFormatterParams{
				StatusCode: http.StatusBadRequest,
			},
			want: yellow,
		},
		{
			name: "Test 5",
			p: &LogFormatterParams{
				StatusCode: http.StatusInternalServerError,
			},
			want: red,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.p.StatusCodeColor(); got != tt.want {
				t.Errorf("StatusCodeColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
