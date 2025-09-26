package pagemeta

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToLowerSlice_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	in := []string{"A", "B", "C"}
	out := toLowerSlice(in)
	if !reflect.DeepEqual(out, []string{"a", "b", "c"}) {
		t.Errorf("toLowerSlice() = %v, want %v", out, []string{"a", "b", "c"})
	}
}
