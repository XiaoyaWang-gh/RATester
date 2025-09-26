package alils

import (
	"fmt"
	"testing"
)

func TestGetAppliedMachineGroups_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	project := &LogProject{
		Name:            "test_project",
		Endpoint:        "http://localhost:8080",
		AccessKeyID:     "test_key_id",
		AccessKeySecret: "test_key_secret",
	}

	confName := "test_config"
	groupNames, err := project.GetAppliedMachineGroups(confName)
	if err != nil {
		t.Errorf("Failed to get applied machine groups: %v", err)
	}

	if len(groupNames) == 0 {
		t.Errorf("No machine groups applied to the config")
	}
}
