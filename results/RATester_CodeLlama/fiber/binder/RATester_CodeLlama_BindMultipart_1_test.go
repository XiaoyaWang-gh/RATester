package binder

import (
	"fmt"
	"testing"
)

func TestBindMultipart_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: add test cases
}
