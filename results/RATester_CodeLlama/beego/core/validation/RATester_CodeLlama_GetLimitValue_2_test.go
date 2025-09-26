package validation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetLimitValue_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := Range{
		Min: Min{
			Min: 1,
		},
		Max: Max{
			Max: 10,
		},
	}
	if !reflect.DeepEqual(r.GetLimitValue(), []int{1, 10}) {
		t.Errorf("GetLimitValue() = %v, want %v", r.GetLimitValue(), []int{1, 10})
	}
}
