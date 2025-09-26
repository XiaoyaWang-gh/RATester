package validation

import (
	"fmt"
	"testing"
)

func TestCheck_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &Validation{}
	obj := &struct {
		Name string
	}{}
	checks := []Validator{
		&Alpha{},
	}
	result := v.Check(obj, checks...)
	if result.Error == nil {
		t.Errorf("Expected error, but got nil")
	}
}
