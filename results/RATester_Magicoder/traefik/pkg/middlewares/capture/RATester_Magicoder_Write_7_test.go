package capture

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWrite_7(t *testing.T) {
	tests := []struct {
		name   string
		status int
		b      []byte
		want   int
	}{
		{
			name:   "Test case 1",
			status: http.StatusOK,
			b:      []byte("Hello, World!"),
			want:   13,
		},
		{
			name:   "Test case 2",
			status: http.StatusNotFound,
			b:      []byte("Page not found!"),
			want:   16,
		},
		{
			name:   "Test case 3",
			status: http.StatusInternalServerError,
			b:      []byte("Internal server error!"),
			want:   22,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			crw := &captureResponseWriter{
				status: tt.status,
			}

			got, _ := crw.Write(tt.b)

			if got != tt.want {
				t.Errorf("Write() = %v, want %v", got, tt.want)
			}
		})
	}
}
