package log

import (
	"fmt"
	"testing"
)

func TestError_2(t *testing.T) {
	t.Parallel()
	t.Run("test case 1", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()
		// Given
		// When
		// Then
	})
	t.Run("test case 2", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()
		// Given
		// When
		// Then
	})
}
