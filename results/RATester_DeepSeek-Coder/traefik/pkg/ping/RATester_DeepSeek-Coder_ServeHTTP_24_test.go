package ping

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_24(t *testing.T) {
	tests := []struct {
		name           string
		terminating    bool
		statusCode     int
		expectedStatus int
	}{
		{
			name:           "Test when handler is not terminating",
			terminating:    false,
			statusCode:     http.StatusOK,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Test when handler is terminating",
			terminating:    true,
			statusCode:     http.StatusServiceUnavailable,
			expectedStatus: http.StatusServiceUnavailable,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, "/", nil)

			handler := &Handler{
				terminating:           tt.terminating,
				TerminatingStatusCode: tt.statusCode,
			}

			handler.ServeHTTP(recorder, request)

			response := recorder.Result()
			defer response.Body.Close()

			if response.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", tt.expectedStatus, response.StatusCode)
			}

			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				t.Errorf("Error reading response body: %v", err)
			}

			if string(body) != http.StatusText(tt.expectedStatus) {
				t.Errorf("Expected body %s, but got %s", http.StatusText(tt.expectedStatus), string(body))
			}
		})
	}
}
