package orm

import (
	"fmt"
	"testing"
)

func TestDistinct_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qs := querySet{}
	qs.Distinct()
	if !qs.distinct {
		t.Error("Distinct method failed")
	}
}
