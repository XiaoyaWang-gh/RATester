package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToStringMapStringE_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	in := map[string]interface{}{
		"a": "1",
		"b": "2",
		"c": "3",
	}
	want := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}
	got, err := ToStringMapStringE(in)
	if err != nil {
		t.Fatalf("ToStringMapStringE() error = %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToStringMapStringE() = %v, want %v", got, want)
	}
}
