package web

import (
	"fmt"
	"testing"
)

func TestStopRun_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic not occurred")
		}
	}()

	ctrl.StopRun()
}
