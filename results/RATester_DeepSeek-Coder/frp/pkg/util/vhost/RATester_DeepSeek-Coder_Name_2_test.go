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

	expected := "test"
	actual := l.Name()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
