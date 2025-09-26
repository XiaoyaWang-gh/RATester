package logs

import (
	"fmt"
	"testing"
)

func TestDestroy_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &consoleWriter{}
	w.Destroy()
}
