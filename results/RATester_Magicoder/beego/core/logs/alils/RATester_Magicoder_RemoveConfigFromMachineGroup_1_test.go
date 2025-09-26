package alils

import (
	"fmt"
	"testing"
)

func TestRemoveConfigFromMachineGroup_1(t *testing.T) {
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
	groupName := "test_group"

	err := project.RemoveConfigFromMachineGroup(confName, groupName)
	if err != nil {
		t.Errorf("Failed to remove config from machine group: %v", err)
	}
}
