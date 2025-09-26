package logs

import (
	"fmt"
	"testing"
)

func TestNeedRotateDaily_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &fileLogWriter{
		MaxLines:         1000,
		maxLinesCurLines: 1001,
		Daily:            true,
		dailyOpenDate:    1,
	}

	if !w.needRotateDaily(2) {
		t.Errorf("Expected true, got false")
	}

	w.maxLinesCurLines = 999
	if w.needRotateDaily(2) {
		t.Errorf("Expected false, got true")
	}

	w.MaxLines = 0
	if w.needRotateDaily(2) {
		t.Errorf("Expected false, got true")
	}

	w.MaxSize = 1000
	w.maxSizeCurSize = 1001
	if !w.needRotateDaily(2) {
		t.Errorf("Expected true, got false")
	}

	w.maxSizeCurSize = 999
	if w.needRotateDaily(2) {
		t.Errorf("Expected false, got true")
	}

	w.MaxSize = 0
	if w.needRotateDaily(2) {
		t.Errorf("Expected false, got true")
	}

	w.Daily = false
	if w.needRotateDaily(2) {
		t.Errorf("Expected false, got true")
	}
}
