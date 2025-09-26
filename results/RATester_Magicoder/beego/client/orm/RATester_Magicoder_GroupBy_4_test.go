package orm

import (
	"fmt"
	"testing"
)

func TestGroupBy_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qs := querySet{}
	qs.GroupBy("id")
	if len(qs.groups) != 1 || qs.groups[0] != "id" {
		t.Errorf("GroupBy failed, expected: [id], got: %v", qs.groups)
	}
}
