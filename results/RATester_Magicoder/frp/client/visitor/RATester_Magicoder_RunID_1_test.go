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

	expected := "test_run_id"
	actual := v.RunID()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
