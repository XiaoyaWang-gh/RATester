package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWrite_5(t *testing.T) {
	testCases := []struct {
		name        string
		modifierErr error
		writeInput  []byte
		expectedErr error
		expectedN   int
	}{
		{
			name:        "Successful write",
			modifierErr: nil,
			writeInput:  []byte("test"),
			expectedErr: nil,
			expectedN:   4,
		},
		{
			name:        "Write error from modifier",
			modifierErr: errors.New("modifier error"),
			writeInput:  []byte("test"),
			expectedErr: errors.New("modifier error"),
			expectedN:   0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			rm := &ResponseModifier{
				req:         req,
				rw:          rr,
				modifierErr: tc.modifierErr,
			}

			n, err := rm.Write(tc.writeInput)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			if n != tc.expectedN {
				t.Errorf("Expected %d bytes written, got %d", tc.expectedN, n)
			}
		})
	}
}
