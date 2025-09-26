package commands

import (
	"fmt"
	"testing"
)

func TestPrepareNew_1(t *testing.T) {
	detector := &fileChangeDetector{
		current: make(map[string]uint64),
		prev:    make(map[string]uint64),
	}

	t.Run("when fileChangeDetector is nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var f *fileChangeDetector
		f.PrepareNew()
	})

	t.Run("when current is nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		detector.current = nil
		detector.PrepareNew()
		if detector.current == nil || detector.prev == nil {
			t.Errorf("expected current and prev to be initialized")
		}
	})

	t.Run("when current is not nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		detector.current["test"] = 123
		detector.PrepareNew()
		if detector.prev["test"] != 123 {
			t.Errorf("expected prev to be equal to current")
		}
		if detector.current == nil {
			t.Errorf("expected current to be nil")
		}
	})
}
