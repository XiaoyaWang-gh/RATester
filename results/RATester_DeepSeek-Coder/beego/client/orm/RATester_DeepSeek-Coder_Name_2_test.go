package orm

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

	d := driver("test")
	expected := "test"
	result := d.Name()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
