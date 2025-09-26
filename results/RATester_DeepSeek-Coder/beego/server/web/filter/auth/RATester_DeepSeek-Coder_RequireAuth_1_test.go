package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequireAuth_1(t *testing.T) {
	testCases := []struct {
		name           string
		realm          string
		expectedHeader string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Test with valid realm",
			realm:          "testRealm",
			expectedHeader: "Basic realm=\"testRealm\"",
			expectedStatus: 401,
			expectedBody:   "401 Unauthorized\n",
		},
		{
			name:           "Test with empty realm",
			realm:          "",
			expectedHeader: "Basic realm=\"\"",
			expectedStatus: 401,
			expectedBody:   "401 Unauthorized\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)

			auth := &BasicAuth{
				Realm: tc.realm,
			}

			auth.RequireAuth(w, r)

			result := w.Result()
			defer result.Body.Close()

			body, _ := ioutil.ReadAll(result.Body)

			if result.Header.Get("WWW-Authenticate") != tc.expectedHeader {
				t.Errorf("Expected header '%s', got '%s'", tc.expectedHeader, result.Header.Get("WWW-Authenticate"))
			}

			if result.StatusCode != tc.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tc.expectedStatus, result.StatusCode)
			}

			if string(body) != tc.expectedBody {
				t.Errorf("Expected body '%s', got '%s'", tc.expectedBody, string(body))
			}
		})
	}
}
