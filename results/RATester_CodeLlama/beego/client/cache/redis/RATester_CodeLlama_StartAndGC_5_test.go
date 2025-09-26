package redis

import (
	"fmt"
	"testing"
)

func TestStartAndGC_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: add your test case here
}
