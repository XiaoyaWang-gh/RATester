package memory

import (
	"fmt"
	"testing"
)

func TestGet_6(t *testing.T) {
	storage := &Storage{
		data: map[string]item{
			"key1": {v: "value1", e: 0},
			"key2": {v: "value2", e: 1},
			"key3": {v: "value3", e: 2},
		},
	}

	t.Run("TestGet", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := storage.Get("key1")
		want := "value1"
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("TestGetExpired", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := storage.Get("key2")
		if got != nil {
			t.Errorf("got %v want nil", got)
		}
	})

	t.Run("TestGetNonExistent", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := storage.Get("key4")
		if got != nil {
			t.Errorf("got %v want nil", got)
		}
	})
}
