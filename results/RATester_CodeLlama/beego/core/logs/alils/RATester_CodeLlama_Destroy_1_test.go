package alils

import (
	"fmt"
	"testing"
)

func TestDestroy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &aliLSWriter{}
	c.Destroy()
}
