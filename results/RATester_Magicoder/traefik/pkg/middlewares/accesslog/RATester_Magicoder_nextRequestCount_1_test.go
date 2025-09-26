package accesslog

import (
	"fmt"
	"testing"
)

func TestNextRequestCount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	initialValue := nextRequestCount()
	nextValue := nextRequestCount()

	if nextValue != initialValue+1 {
		t.Errorf("Expected next value to be %d, but got %d", initialValue+1, nextValue)
	}
}
