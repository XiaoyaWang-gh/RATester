package protoexample

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestGetOptionalgroup_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	test := &Test{
		Label: proto.String("test"),
		Type:  proto.Int32(77),
		Optionalgroup: &Test_OptionalGroup{
			RequiredField: proto.String("required"),
		},
	}

	og := test.GetOptionalgroup()
	if og == nil {
		t.Errorf("Expected non-nil OptionalGroup, got nil")
	}

	if og.RequiredField == nil {
		t.Errorf("Expected non-nil RequiredField, got nil")
	}

	if *og.RequiredField != "required" {
		t.Errorf("Expected RequiredField to be 'required', got %s", *og.RequiredField)
	}
}
