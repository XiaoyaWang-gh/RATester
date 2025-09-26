package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSeq_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("valid arguments", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		seq, err := ns.Seq(1, 5, 1)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		expected := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(seq, expected) {
			t.Errorf("expected %v, got %v", expected, seq)
		}
	})

	t.Run("invalid number of arguments", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Seq()
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("invalid arguments", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Seq("a", "b", "c")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("size limit exceeded", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Seq(1, 100001, 1)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("increment 0", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Seq(1, 5, 0)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("increment direction mismatch", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Seq(5, 1, -1)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
