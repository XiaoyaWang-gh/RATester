package recovery

import (
	"fmt"
	"net/http"
	"testing"
)

func TestShouldLogPanic_1(t *testing.T) {

	tests := []struct {
		name        string
		panicValue  interface{}
		expectedLog bool
	}{
		{
			name:        "Test case 1: panicValue is nil",
			panicValue:  nil,
			expectedLog: false,
		},
		{
			name:        "Test case 2: panicValue is http.ErrAbortHandler",
			panicValue:  http.ErrAbortHandler,
			expectedLog: false,
		},
		{
			name:        "Test case 3: panicValue is not nil or http.ErrAbortHandler",
			panicValue:  "some error",
			expectedLog: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			logged := shouldLogPanic(test.panicValue)
			if logged != test.expectedLog {
				t.Errorf("Expected shouldLogPanic(%v) to be %v, but got %v", test.panicValue, test.expectedLog, logged)
			}
		})
	}
}
