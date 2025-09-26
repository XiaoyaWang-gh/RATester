package failover

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSetFallbackHandler_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc     string
		handler  http.Handler
		expected http.Handler
	}{
		{
			desc:     "set fallback handler",
			handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			expected: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
		},
		{
			desc:     "set fallback handler to nil",
			handler:  nil,
			expected: nil,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			f := &Failover{}
			f.SetFallbackHandler(test.handler)

			if f.fallbackHandler != test.expected {
				t.Errorf("expected %v, got %v", test.expected, f.fallbackHandler)
			}

			if f.fallbackStatus != (test.handler != nil) {
				t.Errorf("expected fallbackStatus to be %v, got %v", test.handler != nil, f.fallbackStatus)
			}
		})
	}
}
