package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStack_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	skip := 1
	indent := "  "

	want := []byte("  at TestStack() [/Users/john/go/src/github.com/john/stack/stack_test.go:11]\n")

	got := Stack(skip, indent)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Stack() = %v, want %v", got, want)
	}
}
