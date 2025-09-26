package commands

import (
	"fmt"
	"testing"
)

func TestSetWasErr_1(t *testing.T) {
	e := &hugoBuilderErrState{}

	t.Run("setWasErr", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		e.setWasErr(true)
		if e.waserr != true {
			t.Errorf("Expected waserr to be true, got %v", e.waserr)
		}
	})

	t.Run("setWasErr", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		e.setWasErr(false)
		if e.waserr != false {
			t.Errorf("Expected waserr to be false, got %v", e.waserr)
		}
	})
}
