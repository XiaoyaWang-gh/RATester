package vhost

import (
	"fmt"
	"testing"
)

func TestName_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Listener{name: "test"}
	if l.Name() != "test" {
		t.Fatal("Name() failed")
	}
}
