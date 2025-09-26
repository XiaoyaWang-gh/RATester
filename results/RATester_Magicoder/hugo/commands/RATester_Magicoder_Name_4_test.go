package commands

import (
	"fmt"
	"testing"
)

func TestName_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cmd := &serverCommand{}
	expected := "server"
	actual := cmd.Name()
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
