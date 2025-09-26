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

	qs := &querySet{
		distinct: false,
	}
	qs.Distinct()
	if !qs.distinct {
		t.Errorf("Distinct() did not set the distinct field to true")
	}
}
