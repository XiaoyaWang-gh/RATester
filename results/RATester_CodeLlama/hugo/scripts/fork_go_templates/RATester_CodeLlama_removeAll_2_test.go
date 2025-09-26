package main

import (
	"fmt"
	"testing"
)

func TestRemoveAll_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expression := "a"
	content := "a"
	want := ""
	if got := removeAll(expression, content); got != want {
		t.Errorf("removeAll() = %v, want %v", got, want)
	}
}
