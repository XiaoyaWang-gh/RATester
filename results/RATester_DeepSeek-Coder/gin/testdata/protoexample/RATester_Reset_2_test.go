package protoexample

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestReset_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testGroup := &Test_OptionalGroup{
		RequiredField: proto.String("test"),
	}

	testGroup.Reset()

	if testGroup.RequiredField != nil {
		t.Errorf("Expected RequiredField to be nil after Reset, got %v", *testGroup.RequiredField)
	}
}
