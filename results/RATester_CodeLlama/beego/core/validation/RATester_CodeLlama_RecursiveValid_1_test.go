package validation

import (
	"fmt"
	"testing"
)

func TestRecursiveValid_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &Validation{}
	objc := &struct {
		Name string `valid:"Required;AlphaNumeric"`
		Age  int    `valid:"Required;Min(1);Max(100)"`
	}{}
	pass, err := v.RecursiveValid(objc)
	if err != nil || !pass {
		t.Errorf("RecursiveValid failed")
	}
}
