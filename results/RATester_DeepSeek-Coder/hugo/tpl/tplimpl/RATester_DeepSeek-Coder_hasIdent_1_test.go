package tplimpl

import (
	"fmt"
	"testing"
)

func TestHasIdent_1(t *testing.T) {
	ctx := &templateContext{}

	t.Run("ReturnsTrueWhenIdentExists", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		idents := []string{"ident1", "ident2", "ident3"}
		ident := "ident2"
		result := ctx.hasIdent(idents, ident)
		if !result {
			t.Errorf("Expected true, got false")
		}
	})

	t.Run("ReturnsFalseWhenIdentDoesNotExist", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		idents := []string{"ident1", "ident2", "ident3"}
		ident := "ident4"
		result := ctx.hasIdent(idents, ident)
		if result {
			t.Errorf("Expected false, got true")
		}
	})
}
