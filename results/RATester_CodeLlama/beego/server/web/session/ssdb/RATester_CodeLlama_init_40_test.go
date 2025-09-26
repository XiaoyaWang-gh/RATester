package ssdb

import (
	"fmt"
	"testing"
)

func TestInit_40(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
