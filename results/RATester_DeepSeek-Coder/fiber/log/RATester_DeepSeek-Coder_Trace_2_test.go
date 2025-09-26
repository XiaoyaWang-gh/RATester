package log

import (
	"fmt"
	"testing"
)

func TestTrace_2(t *testing.T) {
	t.Run("TestTrace", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		Trace("TestTrace")
	})
}
