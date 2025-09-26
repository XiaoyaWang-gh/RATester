package gin

import (
	"fmt"
	"io"
	"testing"
)

func TestStream_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	c := &Context{
		Writer: &responseWriter{},
	}
	c.Stream(func(w io.Writer) bool {
		return true
	})
}
