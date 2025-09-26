package capture

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNeedsReset_1(t *testing.T) {
	capture := &Capture{
		crw: &captureResponseWriter{
			rw:     httptest.NewRecorder(),
			status: http.StatusOK,
			size:   0,
		},
	}

	tests := []struct {
		name string
		rw   http.ResponseWriter
		want bool
	}{
		{
			name: "Same ResponseWriter",
			rw:   capture.crw.rw,
			want: false,
		},
		{
			name: "Different ResponseWriter",
			rw:   httptest.NewRecorder(),
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

			if got := capture.NeedsReset(tt.rw); got != tt.want {
				t.Errorf("NeedsReset() = %v, want %v", got, tt.want)
			}
		})
	}
}
