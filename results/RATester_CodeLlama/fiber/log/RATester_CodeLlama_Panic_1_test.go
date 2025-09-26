package log

import (
	"fmt"
	"testing"
)

func TestPanic_1(t *testing.T) {
	t.Parallel()
	t.Run("Panic", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()
		Panic()
	})
}
