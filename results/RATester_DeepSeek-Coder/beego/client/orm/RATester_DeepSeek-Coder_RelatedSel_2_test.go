package orm

import (
	"fmt"
	"testing"
)

func TestRelatedSel_2(t *testing.T) {
	qs := querySet{
		related: []string{},
	}

	t.Run("test with string param", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		qs.RelatedSel("profile")
		if len(qs.related) != 1 || qs.related[0] != "profile" {
			t.Errorf("RelatedSel with string param failed, expected related: [profile], got: %v", qs.related)
		}
	})

	t.Run("test with int param", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		qs.RelatedSel(2)
		if qs.relDepth != 2 {
			t.Errorf("RelatedSel with int param failed, expected relDepth: 2, got: %v", qs.relDepth)
		}
	})

	t.Run("test with wrong param type", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("RelatedSel with wrong param type did not panic")
			}
		}()
		qs.RelatedSel(1.2)
	})
}
