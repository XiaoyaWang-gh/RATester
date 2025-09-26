package protoexample

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestGetLabel_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	test := &Test{
		Label: proto.String("testLabel"),
	}

	label := test.GetLabel()

	if label != "testLabel" {
		t.Errorf("Expected 'testLabel', got '%s'", label)
	}
}
