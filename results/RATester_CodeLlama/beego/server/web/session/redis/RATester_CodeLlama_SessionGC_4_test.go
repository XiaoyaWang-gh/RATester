package redis

import (
	"fmt"
	"testing"
)

func TestSessionGC_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
