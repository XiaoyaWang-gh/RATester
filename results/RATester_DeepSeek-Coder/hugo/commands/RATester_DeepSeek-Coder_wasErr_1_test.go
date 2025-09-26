package commands

import (
	"fmt"
	"sync"
	"testing"
)

func TestWasErr_1(t *testing.T) {
	e := &hugoBuilderErrState{
		mu:       sync.Mutex{},
		paused:   false,
		builderr: nil,
		waserr:   false,
	}

	t.Run("Testing wasErr() method", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Run("When waserr is false", func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			e.setWasErr(false)
			if e.wasErr() {
				t.Errorf("Expected wasErr() to return false, but got true")
			}
		})

		t.Run("When waserr is true", func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			e.setWasErr(true)
			if !e.wasErr() {
				t.Errorf("Expected wasErr() to return true, but got false")
			}
		})
	})
}
