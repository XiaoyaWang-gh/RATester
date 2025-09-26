package bundler

import (
	"fmt"
	"testing"
)

func TestClose_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &multiReadSeekCloser{}
	r.Close()
}
