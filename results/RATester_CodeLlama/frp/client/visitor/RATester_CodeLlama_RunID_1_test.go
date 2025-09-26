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

	v := &visitorHelperImpl{}
	v.runID = "123456"
	if v.RunID() != "123456" {
		t.Errorf("RunID() = %v, want %v", v.RunID(), "123456")
	}
}
