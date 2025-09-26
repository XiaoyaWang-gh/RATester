package logs

import (
	"fmt"
	"testing"
)

func TestFlush_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	writer := &consoleWriter{}
	writer.Flush()
}
