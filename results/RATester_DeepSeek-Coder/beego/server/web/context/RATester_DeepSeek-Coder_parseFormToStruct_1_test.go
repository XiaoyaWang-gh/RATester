package context

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestParseFormToStruct_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testStruct struct {
		BoolField   bool
		IntField    int
		StringField string
	}

	form := url.Values{
		"BoolField":   []string{"true"},
		"IntField":    []string{"42"},
		"StringField": []string{"test string"},
	}

	objT := reflect.TypeOf(testStruct{})
	objV := reflect.New(objT).Elem()

	err := parseFormToStruct(form, objT, objV)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := testStruct{
		BoolField:   true,
		IntField:    42,
		StringField: "test string",
	}

	if !reflect.DeepEqual(objV.Interface(), expected) {
		t.Errorf("Expected %v, got %v", expected, objV.Interface())
	}
}
