package recovery

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_12(t *testing.T) {
	tests := []struct {
		name           string
		re             *recovery
		rw             *httptest.ResponseRecorder
		req            *http.Request
		wantStatusCode int
	}{
		{
			name: "Test ServeHTTP with panic",
			re: &recovery{
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					panic("test panic")
				}),
			},
			rw:             httptest.NewRecorder(),
			req:            httptest.NewRequest(http.MethodGet, "/", nil),
			wantStatusCode: http.StatusInternalServerError,
		},
		{
			name: "Test ServeHTTP without panic",
			re: &recovery{
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprint(w, "OK")
				}),
			},
			rw:             httptest.NewRecorder(),
			req:            httptest.NewRequest(http.MethodGet, "/", nil),
			wantStatusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.re.ServeHTTP(tt.rw, tt.req)

			if tt.rw.Code != tt.wantStatusCode {
				t.Errorf("Expected status code %d, got %d", tt.wantStatusCode, tt.rw.Code)
			}
		})
	}
}
