package retry

import (
	"fmt"
	"testing"
)

func TestFlush_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &responseWriter{}
	r.Flush()
}
