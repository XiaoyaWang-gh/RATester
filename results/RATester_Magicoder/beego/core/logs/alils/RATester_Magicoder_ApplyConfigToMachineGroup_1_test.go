package alils

import (
	"fmt"
	"testing"
)

func TestApplyConfigToMachineGroup_1(t *testing.T) {
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

	err := project.ApplyConfigToMachineGroup("test_config", "test_group")
	if err != nil {
		t.Errorf("Failed to apply config to machine group: %v", err)
	}
}
