package middleware

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func TestReset_1(t *testing.T) {
	tests := []struct {
		name   string
		reader io.ReadCloser
	}{
		{
			name:   "Test case 1",
			reader: ioutil.NopCloser(strings.NewReader("test")),
		},
		{
			name:   "Test case 2",
			reader: ioutil.NopCloser(strings.NewReader("")),
		},
		{
			name:   "Test case 3",
			reader: ioutil.NopCloser(strings.NewReader("test test test")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			lr := &limitedReader{}
			lr.Reset(tt.reader)

			if lr.reader != tt.reader {
				t.Errorf("Expected reader to be %v, got %v", tt.reader, lr.reader)
			}

			if lr.read != 0 {
				t.Errorf("Expected read to be 0, got %v", lr.read)
			}
		})
	}
}
