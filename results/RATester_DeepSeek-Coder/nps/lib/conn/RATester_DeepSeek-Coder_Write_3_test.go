package conn

import (
	"errors"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/golang/snappy"
)

func TestWrite_3(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected int
		err      error
	}{
		{
			name:     "Successful write",
			input:    []byte("test"),
			expected: 4,
			err:      nil,
		},
		{
			name:     "Failed write",
			input:    []byte(""),
			expected: 0,
			err:      errors.New("write error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &SnappyConn{
				w: snappy.NewBufferedWriter(ioutil.Discard),
				r: snappy.NewReader(nil),
				c: ioutil.NopCloser(nil),
			}

			n, err := s.Write(tc.input)

			if n != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, n)
			}

			if (err != nil && tc.err == nil) || (err == nil && tc.err != nil) {
				t.Errorf("Expected error %v, got %v", tc.err, err)
			}

			if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Errorf("Expected error message %s, got %s", tc.err.Error(), err.Error())
			}
		})
	}
}
