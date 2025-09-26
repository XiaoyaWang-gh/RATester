package protoexample

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestGetType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	test := &Test{
		Type: proto.Int32(1),
	}

	result := test.GetType()

	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}
