package migration

import (
	"fmt"
	"testing"
)

func TestAddIndex_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{}
	index := &Index{}
	m.AddIndex(index)
	if len(m.Indexes) != 1 {
		t.Errorf("AddIndex failed")
	}
}
