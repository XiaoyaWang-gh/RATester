package httpserver

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWithResponse_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		resp     []byte
		expected string
	}{
		{
			name:     "Test case 1",
			resp:     []byte("Hello, world!"),
			expected: "Hello, world!\n",
		},
		{
			name:     "Test case 2",
			resp:     []byte("Goodbye, world!"),
			expected: "Goodbye, world!\n",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			s := &Server{}
			s = WithResponse(tc.resp)(s)

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()

			s.handler.ServeHTTP(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			if string(body) != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, string(body))
			}
		})
	}
}
