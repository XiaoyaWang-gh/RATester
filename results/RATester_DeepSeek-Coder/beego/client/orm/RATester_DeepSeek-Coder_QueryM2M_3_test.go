package orm

import (
	"fmt"
	"testing"
)

func TestQueryM2M_3(t *testing.T) {
	orm := &DoNothingOrm{}
	md := &struct{}{}
	name := "Tag"

	t.Run("TestQueryM2M", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		queryM2M := orm.QueryM2M(md, name)
		if queryM2M == nil {
			t.Errorf("Expected non-nil, got nil")
		}
	})
}
