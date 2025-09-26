package logger

import (
	"fmt"
	"testing"
)

func TestCreateTagMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: Test createTagMap
	t.Log("Test createTagMap")
}
