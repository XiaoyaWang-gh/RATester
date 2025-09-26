package binding

import (
	"fmt"
	"mime/multipart"
	"reflect"
	"testing"
)

func TestSetByMultipartFormFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		Field1 string
		Field2 []string
	}
	var testStruct TestStruct
	var files []*multipart.FileHeader
	isSet, err := setByMultipartFormFile(reflect.ValueOf(&testStruct).Elem(), reflect.StructField{}, files)
	if err != nil {
		t.Errorf("setByMultipartFormFile() error = %v", err)
		return
	}
	if isSet {
		t.Errorf("setByMultipartFormFile() isSet = %v, want %v", isSet, false)
	}
}
