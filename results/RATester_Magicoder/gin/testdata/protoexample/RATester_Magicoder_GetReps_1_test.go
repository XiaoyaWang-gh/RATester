package protoexample

import (
	"fmt"
	"testing"
)

func TestGetReps_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	test := &Test{
		Reps: []int64{1, 2, 3},
	}

	result := test.GetReps()

	if len(result) != len(test.Reps) {
		t.Errorf("Expected %d, got %d", len(test.Reps), len(result))
	}

	for i, v := range result {
		if v != test.Reps[i] {
			t.Errorf("Expected %d, got %d", test.Reps[i], v)
		}
	}
}
