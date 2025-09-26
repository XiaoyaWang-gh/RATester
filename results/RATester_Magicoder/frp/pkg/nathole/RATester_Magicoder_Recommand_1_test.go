package nathole

import (
	"fmt"
	"testing"
)

func TestRecommand_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mhr := &MakeHoleRecords{
		scores: []*BehaviorScore{
			{Mode: 1, Index: 1, Score: 10},
			{Mode: 2, Index: 2, Score: 9},
			{Mode: 3, Index: 3, Score: 8},
		},
	}
	mode, index := mhr.Recommand()
	if mode != 1 || index != 1 {
		t.Errorf("Expected mode 1 and index 1, but got mode %d and index %d", mode, index)
	}
	if mhr.scores[0].Score != 9 {
		t.Errorf("Expected score 9, but got %d", mhr.scores[0].Score)
	}
	if mhr.LastUpdateTime.IsZero() {
		t.Errorf("Expected LastUpdateTime to be non-zero")
	}
}
