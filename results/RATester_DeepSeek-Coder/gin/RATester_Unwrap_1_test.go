package gin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUnwrap_1(t *testing.T) {
	testCases := []struct {
		name string
		w    *responseWriter
		want http.ResponseWriter
	}{
		{
			name: "TestUnwrap",
			w: &responseWriter{
				ResponseWriter: httptest.NewRecorder(),
				size:           10,
				status:         200,
			},
			want: httptest.NewRecorder(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tc.w.Unwrap()
			if got != tc.want {
				t.Errorf("Unwrap() = %v, want %v", got, tc.want)
			}
		})
	}
}
