package protoexample

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestDescriptor_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	test := &Test_OptionalGroup{
		RequiredField: proto.String("test"),
	}

	desc, path := test.Descriptor()

	if len(desc) == 0 {
		t.Error("Descriptor should not be empty")
	}

	if len(path) == 0 {
		t.Error("Path should not be empty")
	}
}
