package redis_sentinel

import (
	"fmt"
	"testing"
)

func TestSessionDestroy_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
