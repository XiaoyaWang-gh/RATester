package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToStringMapE_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	in := Params{
		"a": "1",
		"b": "2",
	}
	want := map[string]any{
		"a": "1",
		"b": "2",
	}
	got, err := ToStringMapE(in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
