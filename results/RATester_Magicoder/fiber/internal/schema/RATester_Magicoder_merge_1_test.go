package schema

import (
	"errors"
	"fmt"
	"testing"
)

func Testmerge_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e1 := MultiError{"key1": errors.New("error1"), "key2": errors.New("error2")}
	e2 := MultiError{"key1": errors.New("error3"), "key3": errors.New("error4")}

	e1.merge(e2)

	if len(e1) != 3 {
		t.Errorf("Expected 3 errors, got %d", len(e1))
	}

	if e1["key1"].Error() != "error1" {
		t.Errorf("Expected error1, got %s", e1["key1"].Error())
	}

	if e1["key2"].Error() != "error2" {
		t.Errorf("Expected error2, got %s", e1["key2"].Error())
	}

	if e1["key3"].Error() != "error4" {
		t.Errorf("Expected error4, got %s", e1["key3"].Error())
	}
}
