package failover

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_25(t *testing.T) {
	tests := []struct {
		name           string
		handlerStatus  bool
		fallbackStatus bool
		expectedStatus int
	}{
		{
			name:           "handler is up",
			handlerStatus:  true,
			fallbackStatus: false,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "fallback is up",
			handlerStatus:  false,
			fallbackStatus: true,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "both are down",
			handlerStatus:  false,
			fallbackStatus: false,
			expectedStatus: http.StatusServiceUnavailable,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			failover := &Failover{
				handlerStatus:  test.handlerStatus,
				fallbackStatus: test.fallbackStatus,
			}

			recorder := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			failover.ServeHTTP(recorder, req)

			if recorder.Code != test.expectedStatus {
				t.Errorf("expected status %d, got %d", test.expectedStatus, recorder.Code)
			}
		})
	}
}
