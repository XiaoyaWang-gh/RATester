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
		records: make(map[string]*MakeHoleRecords),
	}

	key := "127.0.0.1:8080"
	mode := 1
	index := 0

	a.ReportSuccess(key, mode, index)

	if _, ok := a.records[key]; !ok {
		t.Errorf("Expected key %s to be in records", key)
	}
}
