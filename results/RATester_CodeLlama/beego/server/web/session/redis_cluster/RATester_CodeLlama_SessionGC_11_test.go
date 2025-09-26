package redis_cluster

import (
	"fmt"
	"testing"
)

func TestSessionGC_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
