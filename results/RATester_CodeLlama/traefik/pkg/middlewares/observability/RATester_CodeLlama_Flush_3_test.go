package observability

import (
	"fmt"
	"testing"
)

func TestFlush_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &statusCodeRecorder{}
	s.Flush()
}
