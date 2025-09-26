package testenv

import (
	"fmt"
	"testing"
)

func TestMustHaveBuildMode_1(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		MustHaveBuildMode(t, "debug")
		if !t.Failed() {
			t.Errorf("Expected test to fail due to missing build mode")
		}
	})

	t.Run("test case 2", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		MustHaveBuildMode(t, "release")
		if t.Failed() {
			t.Errorf("Expected test not to fail due to valid build mode")
		}
	})
}
