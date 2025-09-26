package hugo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrev_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := Version{Major: 1, Minor: 2}
	want := Version{Major: 1, Minor: 1}
	if got := v.Prev(); !reflect.DeepEqual(got, want) {
		t.Errorf("Prev() = %v, want %v", got, want)
	}
}
