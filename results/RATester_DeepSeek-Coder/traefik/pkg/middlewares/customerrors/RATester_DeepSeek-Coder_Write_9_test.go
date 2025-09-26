package customerrors

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWrite_9(t *testing.T) {
	testCases := []struct {
		name     string
		code     int
		buf      []byte
		expected int
	}{
		{
			name:     "Test case 1",
			code:     200,
			buf:      []byte("Hello, World!"),
			expected: 13,
		},
		{
			name:     "Test case 2",
			code:     404,
			buf:      []byte("Not Found"),
			expected: 9,
		},
		{
			name:     "Test case 3",
			code:     500,
			buf:      []byte("Internal Server Error"),
			expected: 19,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			rw := httptest.NewRecorder()
			cm := &codeModifier{
				code:           tc.code,
				headerSent:     false,
				headerMap:      make(http.Header),
				responseWriter: rw,
			}

			n, err := cm.Write(tc.buf)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if n != tc.expected {
				t.Errorf("Expected %d bytes written, but got %d", tc.expected, n)
			}

			if rw.Code != tc.code {
				t.Errorf("Expected response code %d, but got %d", tc.code, rw.Code)
			}

			if rw.Body.String() != string(tc.buf) {
				t.Errorf("Expected body %q, but got %q", string(tc.buf), rw.Body.String())
			}
		})
	}
}
