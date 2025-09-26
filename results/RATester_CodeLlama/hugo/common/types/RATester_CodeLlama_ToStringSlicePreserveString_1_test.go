package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToStringSlicePreserveString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := []string{"a", "b", "c"}
	vv := ToStringSlicePreserveString(v)
	if !reflect.DeepEqual(vv, v) {
		t.Errorf("ToStringSlicePreserveString() = %v, want %v", vv, v)
	}
}
