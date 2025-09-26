package protoexample

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestProtoMessage_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	test := &Test{
		Label: proto.String("test"),
		Type:  proto.Int32(77),
		Reps:  []int64{1, 2, 3},
	}

	// Test Label
	if test.GetLabel() != "test" {
		t.Errorf("Expected Label to be 'test', got %s", test.GetLabel())
	}

	// Test Type
	if *test.Type != 77 {
		t.Errorf("Expected Type to be 77, got %d", *test.Type)
	}

	// Test Reps
	if len(test.Reps) != 3 {
		t.Errorf("Expected 3 Reps, got %d", len(test.Reps))
	}

	for i, rep := range test.Reps {
		if rep != int64(i+1) {
			t.Errorf("Expected Rep %d to be %d, got %d", i+1, i+1, rep)
		}
	}
}
