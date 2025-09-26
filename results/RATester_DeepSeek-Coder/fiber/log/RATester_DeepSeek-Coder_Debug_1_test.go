package log

import (
	"fmt"
	"testing"
)

func TestDebug_1(t *testing.T) {
	t.Run("TestDebug", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		Debug("TestDebug")
	})
}
