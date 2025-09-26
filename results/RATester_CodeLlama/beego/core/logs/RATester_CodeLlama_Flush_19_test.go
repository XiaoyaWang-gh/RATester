package logs

import (
	"fmt"
	"testing"
)

func TestFlush_19(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &SLACKWriter{}
	s.Flush()
}
