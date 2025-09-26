package hugolib

import (
	"fmt"
	"testing"
)

func TestKey_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageState{
		pid: 123,
	}
	expected := "page-123"
	actual := p.Key()
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
