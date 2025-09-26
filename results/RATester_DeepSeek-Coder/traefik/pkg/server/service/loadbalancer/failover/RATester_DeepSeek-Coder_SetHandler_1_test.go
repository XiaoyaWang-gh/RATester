package failover

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSetHandler_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc     string
		handler  http.Handler
		expected http.Handler
	}{
		{
			desc:     "set handler",
			handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			expected: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
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
			f.SetHandler(test.handler)

			if f.handler != test.expected {
				t.Errorf("expected %v, got %v", test.expected, f.handler)
			}

			if f.handlerStatus != true {
				t.Errorf("expected handler status to be true, got %v", f.handlerStatus)
			}
		})
	}
}
