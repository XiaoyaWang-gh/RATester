package context

import (
	"fmt"
	"io"
	"testing"
)

func TestReset_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var n nopResetWriter
	var w io.Writer
	n.Reset(w)
}
