package context

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestBindStruct_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		Name string
		Age  int
	}
	input := &BeegoInput{
		pnames:  []string{"Name", "Age"},
		pvalues: []string{"test", "18"},
	}
	typ := reflect.TypeOf(TestStruct{})
	result := input.bindStruct(&url.Values{}, "test", typ)
	if result.FieldByName("Name").String() != "test" {
		t.Errorf("TestbindStruct failed")
	}
	if result.FieldByName("Age").Int() != 18 {
		t.Errorf("TestbindStruct failed")
	}
}
