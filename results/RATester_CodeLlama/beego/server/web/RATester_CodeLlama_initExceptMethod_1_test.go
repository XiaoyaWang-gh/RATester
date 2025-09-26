package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInitExceptMethod_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	want := []string{"initExceptMethod"}
	got := initExceptMethod()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
