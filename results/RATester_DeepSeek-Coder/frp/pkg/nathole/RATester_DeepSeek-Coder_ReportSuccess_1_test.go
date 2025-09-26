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

	key := "192.168.1.1:1234"
	mode := 0
	index := 1

	a.ReportSuccess(key, mode, index)

	a.mu.Lock()
	records, ok := a.records[key]
	a.mu.Unlock()

	if !ok {
		t.Errorf("Expected key %s to exist in records", key)
	}

	records.ReportSuccess(mode, index)
}
