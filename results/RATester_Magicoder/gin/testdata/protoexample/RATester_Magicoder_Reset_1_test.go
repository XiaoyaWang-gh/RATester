package protoexample

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestReset_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	test := &Test{
		Label: proto.String("test"),
		Type:  proto.Int32(1),
		Reps:  []int64{1, 2, 3},
	}

	test.Reset()

	if test.Label != nil {
		t.Errorf("Label should be nil after Reset, but got %v", *test.Label)
	}

	if test.Type != nil {
		t.Errorf("Type should be nil after Reset, but got %v", *test.Type)
	}

	if len(test.Reps) != 0 {
		t.Errorf("Reps should be empty after Reset, but got %v", test.Reps)
	}
}
