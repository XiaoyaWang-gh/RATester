package redis

import (
	"fmt"
	"testing"
)

func TestGet_38(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
