package cache

import (
	"fmt"
	"testing"
)

func TestInit_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	Register("memory", NewMemoryCache)
}
