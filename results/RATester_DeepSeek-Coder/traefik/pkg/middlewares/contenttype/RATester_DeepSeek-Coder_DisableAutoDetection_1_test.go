package contenttype

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestDisableAutoDetection_1(t *testing.T) {
	testCases := []struct {
		name           string
		handler        http.Handler
		expectedHeader http.Header
	}{
		{
			name: "Test with existing Content-Type header",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
			}),
			expectedHeader: http.Header{"Content-Type": []string{"application/json"}},
		},
		{
			name: "Test with no Content-Type header",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// No Content-Type header set
			}),
			expectedHeader: http.Header{"Content-Type": nil},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			rr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/", nil)

			DisableAutoDetection(tc.handler).ServeHTTP(rr, req)

			if !reflect.DeepEqual(rr.Header(), tc.expectedHeader) {
				t.Errorf("Expected header %v, but got %v", tc.expectedHeader, rr.Header())
			}
		})
	}
}
