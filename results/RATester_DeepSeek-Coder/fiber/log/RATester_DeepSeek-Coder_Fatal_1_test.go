package log

import (
	"fmt"
	"testing"
)

func TestFatal_1(t *testing.T) {
	t.Run("TestFatal", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		Fatal("TestFatal")
	})
}
