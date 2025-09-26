package file

import (
	"fmt"
	"testing"
)

func TestLen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := PairList{
		&Pair{key: "a", cId: 1, order: "asc", clientFlow: &Flow{ExportFlow: 100}},
		&Pair{key: "b", cId: 2, order: "desc", clientFlow: &Flow{ExportFlow: 200}},
		&Pair{key: "c", cId: 3, order: "asc", clientFlow: &Flow{ExportFlow: 300}},
	}
	if p.Len() != 3 {
		t.Errorf("p.Len() = %d, want 3", p.Len())
	}
}
