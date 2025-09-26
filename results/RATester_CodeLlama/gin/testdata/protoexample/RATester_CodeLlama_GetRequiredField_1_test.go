package protoexample

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestGetRequiredField_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var test Test_OptionalGroup
	test.RequiredField = proto.String("test")
	if test.GetRequiredField() != "test" {
		t.Errorf("GetRequiredField() = %v, want %v", test.GetRequiredField(), "test")
	}
}
