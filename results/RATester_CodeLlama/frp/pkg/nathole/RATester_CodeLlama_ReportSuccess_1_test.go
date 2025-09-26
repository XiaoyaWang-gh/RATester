package nathole

import (
	"fmt"
	"testing"
)

func TestReportSuccess_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := &Analyzer{
		records: map[string]*MakeHoleRecords{
			"key": {
				scores: []*BehaviorScore{
					{Mode: 1, Index: 1, Score: 10},
					{Mode: 1, Index: 2, Score: 10},
					{Mode: 1, Index: 3, Score: 10},
				},
			},
		},
	}
	a.ReportSuccess("key", 1, 1)
	if len(a.records["key"].scores) != 3 {
		t.Fatal("ReportSuccess failed")
	}
	if a.records["key"].scores[0].Score != 10 {
		t.Fatal("ReportSuccess failed")
	}
	if a.records["key"].scores[1].Score != 10 {
		t.Fatal("ReportSuccess failed")
	}
	if a.records["key"].scores[2].Score != 10 {
		t.Fatal("ReportSuccess failed")
	}
}
