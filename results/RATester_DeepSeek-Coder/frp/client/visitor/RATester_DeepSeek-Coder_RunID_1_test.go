package visitor

import (
	"fmt"
	"testing"
)

func TestRunID_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &visitorHelperImpl{
		runID: "test_run_id",
	}

	runID := v.RunID()
	if runID != "test_run_id" {
		t.Errorf("Expected runID to be 'test_run_id', but got '%s'", runID)
	}
}
