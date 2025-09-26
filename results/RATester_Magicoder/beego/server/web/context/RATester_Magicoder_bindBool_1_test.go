package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestbindBool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{}
	tests := []struct {
		val  string
		want bool
	}{
		{"true", true},
		{"on", true},
		{"1", true},
		{"false", false},
		{"off", false},
		{"0", false},
	}
	for _, test := range tests {
		got := input.bindBool(test.val, reflect.TypeOf(test.want))
		if got.Bool() != test.want {
			t.Errorf("bindBool(%q) = %v, want %v", test.val, got.Bool(), test.want)
		}
	}
}
