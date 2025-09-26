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

	l := &Listener{
		name: "test",
	}

	name := l.Name()

	if name != "test" {
		t.Errorf("Expected 'test', but got '%s'", name)
	}
}
