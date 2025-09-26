package echo

import (
	"fmt"
	"testing"
)

func TestMustUints_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValuesFunc: func(sourceParam string) []string {
			if sourceParam == "ids" {
				return []string{"1", "2", "3"}
			}
			return nil
		},
	}

	var ids []uint
	b.MustUints("ids", &ids)

	if len(ids) != 3 {
		t.Errorf("Expected 3 ids, got %d", len(ids))
	}

	if ids[0] != 1 || ids[1] != 2 || ids[2] != 3 {
		t.Errorf("Expected ids to be [1, 2, 3], got %v", ids)
	}
}
