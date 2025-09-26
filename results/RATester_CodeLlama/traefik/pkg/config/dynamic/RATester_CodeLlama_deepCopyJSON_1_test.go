package dynamic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeepCopyJSON_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	x := map[string]any{
		"a": 1,
		"b": "2",
		"c": []any{3, 4},
	}
	y := deepCopyJSON(x)
	if !reflect.DeepEqual(x, y) {
		t.Errorf("deepCopyJSON(%v) = %v, want %v", x, y, x)
	}
}
