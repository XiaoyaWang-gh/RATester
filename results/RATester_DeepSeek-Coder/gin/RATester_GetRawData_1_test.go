package gin

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestGetRawData_1(t *testing.T) {
	testCases := []struct {
		name     string
		body     io.ReadCloser
		expected []byte
		err      error
	}{
		{
			name:     "Nil body",
			body:     nil,
			expected: nil,
			err:      errors.New("cannot read nil body"),
		},
		{
			name:     "Non-nil body",
			body:     io.NopCloser(strings.NewReader("test")),
			expected: []byte("test"),
			err:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := &Context{
				Request: &http.Request{
					Body: tc.body,
				},
			}

			actual, err := ctx.GetRawData()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
			if !errors.Is(err, tc.err) {
				t.Errorf("Expected error %v, got %v", tc.err, err)
			}
		})
	}
}
