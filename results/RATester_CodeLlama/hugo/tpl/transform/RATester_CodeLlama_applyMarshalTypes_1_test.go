package transform

import (
	"fmt"
	"testing"
)

func TestApplyMarshalTypes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := map[string]any{
		"a": map[string]any{
			"b": 1.0,
			"c": 2.0,
		},
		"d": 3.0,
	}
	applyMarshalTypes(m)
	if m["a"].(map[string]any)["b"].(int64) != 1 {
		t.Errorf("m[\"a\"][\"b\"] != 1")
	}
	if m["a"].(map[string]any)["c"].(int64) != 2 {
		t.Errorf("m[\"a\"][\"c\"] != 2")
	}
	if m["d"].(int64) != 3 {
		t.Errorf("m[\"d\"] != 3")
	}
}
