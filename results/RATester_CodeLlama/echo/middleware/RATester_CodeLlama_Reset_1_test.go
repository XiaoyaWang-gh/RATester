package middleware

import (
	"fmt"
	"io"
	"testing"
)

func TestReset_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var r *limitedReader
	var reader io.ReadCloser
	r.Reset(reader)
}
