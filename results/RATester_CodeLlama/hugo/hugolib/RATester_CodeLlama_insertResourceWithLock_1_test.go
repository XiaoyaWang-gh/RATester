package hugolib

import (
	"fmt"
	"testing"
)

func TestInsertResourceWithLock_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: Add test cases.
}
