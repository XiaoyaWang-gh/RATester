package redis_sentinel

import (
	"fmt"
	"testing"
)

func TestGet_39(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
