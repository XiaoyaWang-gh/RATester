package cache

import (
	"fmt"
	"testing"
)

func TestNewCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: add test cases.
	t.Errorf("Failed to generate test case for NewCache")
}
