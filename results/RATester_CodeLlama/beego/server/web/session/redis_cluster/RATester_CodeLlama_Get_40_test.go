package redis_cluster

import (
	"fmt"
	"testing"
)

func TestGet_40(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
