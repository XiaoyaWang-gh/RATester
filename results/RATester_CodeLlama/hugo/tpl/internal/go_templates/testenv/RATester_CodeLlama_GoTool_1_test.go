package testenv

import (
	"fmt"
	"testing"
)

func TestGoTool_1(t *testing.T) {
	t.Run("test go tool", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		goToolPath, err := GoTool()
		if err != nil {
			t.Errorf("go tool error: %v", err)
		}
		if goToolPath == "" {
			t.Errorf("go tool path is empty")
		}
	})
}
