package resources

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

	l := &genericResource{}
	l.name = "test"
	if l.Name() != "test" {
		t.Errorf("Expected name to be test, got %s", l.Name())
	}
}
