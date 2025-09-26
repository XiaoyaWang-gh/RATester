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
		if t.Failed() {
			t.Errorf("Expected build mode to be 'debug', but it was not.")
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
			t.Errorf("Expected build mode to be 'release', but it was not.")
		}
	})

	t.Run("test case 3", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		MustHaveBuildMode(t, "invalid")
		if !t.Failed() {
			t.Errorf("Expected build mode to be invalid, but it was valid.")
		}
	})
}
