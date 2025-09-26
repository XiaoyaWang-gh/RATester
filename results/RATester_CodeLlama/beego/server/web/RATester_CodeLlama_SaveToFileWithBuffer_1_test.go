package web

import (
	"fmt"
	"testing"
)

func TestSaveToFileWithBuffer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
