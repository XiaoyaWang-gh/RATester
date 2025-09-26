package cache

import (
	"fmt"
	"testing"
)

func TestGetMulti_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
