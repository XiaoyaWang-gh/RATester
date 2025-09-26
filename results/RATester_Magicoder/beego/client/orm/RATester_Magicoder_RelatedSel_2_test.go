package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestRelatedSel_2(t *testing.T) {
	o := querySet{
		mi: &models.ModelInfo{},
	}

	t.Run("Test with no params", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		o.RelatedSel()
		if o.relDepth != DefaultRelsDepth {
			t.Errorf("Expected relDepth to be %d, but got %d", DefaultRelsDepth, o.relDepth)
		}
	})

	t.Run("Test with string param", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		o.RelatedSel("profile")
		if len(o.related) != 1 || o.related[0] != "profile" {
			t.Errorf("Expected related to be [profile], but got %v", o.related)
		}
	})

	t.Run("Test with int param", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		o.RelatedSel(2)
		if o.relDepth != 2 {
			t.Errorf("Expected relDepth to be 2, but got %d", o.relDepth)
		}
	})

	t.Run("Test with wrong param type", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic, but no panic occurred")
			}
		}()
		o.RelatedSel(true)
	})
}
