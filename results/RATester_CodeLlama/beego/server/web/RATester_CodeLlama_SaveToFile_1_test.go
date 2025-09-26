package web

import (
	"fmt"
	"testing"
)

func TestSaveToFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup test
	// TODO: setup mock
	// TODO: run test
	// TODO: assert
}
