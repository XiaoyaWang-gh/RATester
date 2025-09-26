package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetParams_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dst := Params{"a": "a", "b": "b", "c": "c"}
	src := Params{"a": "aa", "b": "bb", "d": "d"}
	SetParams(dst, src)
	if !reflect.DeepEqual(dst, Params{"a": "aa", "b": "bb", "c": "c", "d": "d"}) {
		t.Errorf("dst = %v, want %v", dst, Params{"a": "aa", "b": "bb", "c": "c", "d": "d"})
	}
}
