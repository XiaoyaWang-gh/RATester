package context

import (
	"fmt"
	"testing"
)

func TestFlush_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Response{}
	r.Flush()
}
