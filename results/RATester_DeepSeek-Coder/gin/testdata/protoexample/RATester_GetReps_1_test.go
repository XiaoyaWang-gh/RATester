package protoexample

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestGetReps_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	test := &Test{
		Label: proto.String("test"),
		Reps:  []int64{1, 2, 3},
	}

	reps := test.GetReps()

	if len(reps) != 3 {
		t.Errorf("Expected 3 reps, got %d", len(reps))
	}

	if reps[0] != 1 || reps[1] != 2 || reps[2] != 3 {
		t.Errorf("Expected reps to be [1, 2, 3], got %v", reps)
	}
}
