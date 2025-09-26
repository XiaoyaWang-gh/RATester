package testenv

import (
	"fmt"
	"testing"
)

func TestMustHaveBuildMode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Helper()
	t.Cleanup(func() {
		if r := recover(); r != nil {
			t.Errorf("panic: %v", r)
		}
	})
	MustHaveBuildMode(t, "buildmode")
}
