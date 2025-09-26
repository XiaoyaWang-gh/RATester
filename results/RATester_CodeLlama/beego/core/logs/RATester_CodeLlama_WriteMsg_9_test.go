package logs

import (
	"fmt"
	"testing"
)

func TestWriteMsg_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
