package memcache

import (
	"fmt"
	"testing"
)

func TestClearAll_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
