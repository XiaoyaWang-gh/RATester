package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBindPoint_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{}
	key := "key"
	typ := reflect.TypeOf(1)
	got := input.bindPoint(key, typ)
	want := reflect.ValueOf(1).Addr()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("input.bindPoint() = %v, want %v", got, want)
	}
}
