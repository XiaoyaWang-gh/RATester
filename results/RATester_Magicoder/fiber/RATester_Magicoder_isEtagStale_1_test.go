package fiber

import (
	"fmt"
	"testing"
)

func TestisEtagStale_1(t *testing.T) {
	app := &App{
		getString: func(b []byte) string {
			return string(b)
		},
	}

	tests := []struct {
		name           string
		etag           string
		noneMatchBytes []byte
		want           bool
	}{
		{
			name:           "Etag is stale",
			etag:           "123",
			noneMatchBytes: []byte("\"123\", \"456\""),
			want:           true,
		},
		{
			name:           "Etag is not stale",
			etag:           "123",
			noneMatchBytes: []byte("\"456\""),
			want:           false,
		},
		{
			name:           "Etag is not stale with multiple commas",
			etag:           "123",
			noneMatchBytes: []byte("\"456\", \"789\", \"123\""),
			want:           false,
		},
		{
			name:           "Etag is stale with multiple commas",
			etag:           "123",
			noneMatchBytes: []byte("\"456\", \"789\", \"123\", \"456\""),
			want:           true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := app.isEtagStale(tt.etag, tt.noneMatchBytes); got != tt.want {
				t.Errorf("isEtagStale() = %v, want %v", got, tt.want)
			}
		})
	}
}
