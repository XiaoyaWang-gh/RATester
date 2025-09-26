package common

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTrimArr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	arr := []string{"", "a", "b", "c", "", "d", "e", "", "f"}
	want := []string{"a", "b", "c", "d", "e", "f"}
	got := TrimArr(arr)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TrimArr() = %v, want %v", got, want)
	}
}
