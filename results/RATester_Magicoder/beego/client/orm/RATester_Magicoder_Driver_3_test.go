package orm

import (
	"fmt"
	"testing"
)

func TestDriver_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &ormBase{
		alias: &alias{
			Name: "test",
		},
	}

	expected := driver("test")
	actual := o.Driver()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
