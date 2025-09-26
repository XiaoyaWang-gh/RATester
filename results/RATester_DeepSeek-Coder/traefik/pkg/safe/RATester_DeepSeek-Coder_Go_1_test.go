package safe

import (
	"fmt"
	"testing"
)

func TestGo_1(t *testing.T) {
	t.Run("should start a goroutine and recover from panic", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		Go(func() {
			panic("Panic in goroutine")
		})
	})
}
