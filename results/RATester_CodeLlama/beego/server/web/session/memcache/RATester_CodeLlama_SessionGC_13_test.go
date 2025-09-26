package memcache

import (
	"fmt"
	"testing"
)

func TestSessionGC_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
