package binding

import (
	"fmt"
	"mime/multipart"
	"reflect"
	"testing"
)

func TestSetArrayOfMultipartFormFiles_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type Test struct {
		Name string
		Age  int
	}
	var test Test
	var files []*multipart.FileHeader
	isSet, err := setArrayOfMultipartFormFiles(reflect.ValueOf(&test).Elem(), reflect.TypeOf(test).Field(0), files)
	if err != nil {
		t.Errorf("setArrayOfMultipartFormFiles() error = %v", err)
		return
	}
	if isSet {
		t.Errorf("setArrayOfMultipartFormFiles() isSet = %v, want %v", isSet, false)
	}
}
