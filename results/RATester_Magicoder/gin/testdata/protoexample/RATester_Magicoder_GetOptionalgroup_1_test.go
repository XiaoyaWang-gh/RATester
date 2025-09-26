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
		Type:  proto.Int32(1),
		Reps:  []int64{1, 2, 3},
		Optionalgroup: &Test_OptionalGroup{
			RequiredField: proto.String("required"),
		},
	}

	result := test.GetOptionalgroup()

	if result == nil {
		t.Error("Expected non-nil result, but got nil")
	}

	if result.RequiredField == nil {
		t.Error("Expected non-nil RequiredField, but got nil")
	}

	if *result.RequiredField != "required" {
		t.Errorf("Expected RequiredField to be 'required', but got '%s'", *result.RequiredField)
	}
}
