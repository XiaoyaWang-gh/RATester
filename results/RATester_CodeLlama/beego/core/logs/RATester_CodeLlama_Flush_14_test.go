package logs

import (
	"fmt"
	"testing"
)

func TestFlush_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &connWriter{}
	c.Flush()
}
