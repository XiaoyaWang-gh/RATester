package pageparser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInput_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{}
	l.input = []byte("input")
	if got := l.Input(); !reflect.DeepEqual(got, l.input) {
		t.Errorf("Input() = %v, want %v", got, l.input)
	}
}
