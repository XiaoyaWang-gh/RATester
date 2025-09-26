package strings

import (
	"fmt"
	"testing"
)

func TestFindRE_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		re, err := ns.FindRE("a(x*)b", "axxb")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(re) != 1 || re[0] != "axxb" {
			t.Errorf("Expected ['axxb'], got %v", re)
		}
	})

	t.Run("error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.FindRE("(", "test")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("limit", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		re, err := ns.FindRE("a", "aaaaa", 3)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(re) != 3 || re[0] != "a" || re[1] != "a" || re[2] != "a" {
			t.Errorf("Expected ['a', 'a', 'a'], got %v", re)
		}
	})

	t.Run("non-string content", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.FindRE("a", 123)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("non-string limit", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.FindRE("a", "test", "b")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
