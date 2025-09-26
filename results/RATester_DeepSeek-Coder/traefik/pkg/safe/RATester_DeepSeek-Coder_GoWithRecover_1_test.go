package safe

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestGoWithRecover_1(t *testing.T) {
	t.Run("Test panic handling", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testErr := errors.New("test error")
		testPanic := func() {
			panic(testErr)
		}
		recovered := false
		testRecover := func(err interface{}) {
			if err == testErr {
				recovered = true
			}
		}
		GoWithRecover(testPanic, testRecover)
		time.Sleep(1 * time.Second) // wait for goroutine to finish
		if !recovered {
			t.Errorf("Expected to recover from panic, but did not")
		}
	})

	t.Run("Test no panic", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testNoPanic := func() {
			// do nothing
		}
		recovered := false
		testRecover := func(err interface{}) {
			recovered = true
		}
		GoWithRecover(testNoPanic, testRecover)
		time.Sleep(1 * time.Second) // wait for goroutine to finish
		if recovered {
			t.Errorf("Expected not to recover from no panic, but did")
		}
	})
}
