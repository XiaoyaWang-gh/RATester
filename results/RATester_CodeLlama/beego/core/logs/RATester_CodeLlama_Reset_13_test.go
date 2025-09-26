package logs

import (
	"fmt"
	"testing"
)

func TestReset_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{}
	bl.Reset()
}
