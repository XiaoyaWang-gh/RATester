package alils

import (
	"fmt"
	"testing"
)

func TestGetAppliedMachineGroup_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup test
	c := &LogConfig{}
	confName := "test"
	groupNames, err := c.GetAppliedMachineGroup(confName)
	if err != nil {
		t.Errorf("GetAppliedMachineGroup() error = %v", err)
		return
	}
	if len(groupNames) != 0 {
		t.Errorf("GetAppliedMachineGroup() groupNames = %v, want %v", groupNames, 0)
	}
	// TODO: teardown test
}
