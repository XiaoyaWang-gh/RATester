package binding

import (
	"fmt"
	"testing"
)

func TestMapFormByTag_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type Test struct {
		Name string `form:"name"`
	}
	var test Test
	form := map[string][]string{
		"name": {"test"},
	}
	err := mapFormByTag(&test, form, "form")
	if err != nil {
		t.Errorf("mapFormByTag error: %v", err)
	}
	if test.Name != "test" {
		t.Errorf("mapFormByTag error: %v", test.Name)
	}
}
