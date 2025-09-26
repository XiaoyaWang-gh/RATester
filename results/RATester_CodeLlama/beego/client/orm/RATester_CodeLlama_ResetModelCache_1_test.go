package orm

import (
	"fmt"
	"testing"
)

func TestResetModelCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	// when
	ResetModelCache()
	// then
	// ...
}
