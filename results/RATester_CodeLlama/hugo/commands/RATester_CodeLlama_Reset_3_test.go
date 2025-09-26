package commands

import (
	"fmt"
	"testing"
)

func TestReset_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &rootCommand{}
	r.Reset()
}
