package alils

import (
	"fmt"
	"testing"
)

func TestUpdateMachineGroup_1(t *testing.T) {
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

	machineGroup := &MachineGroup{
		Name:          "test_group",
		Type:          "test_type",
		MachineIDType: "test_id_type",
		MachineIDList: []string{"test_machine_id1", "test_machine_id2"},
		project:       project,
	}

	err := project.UpdateMachineGroup(machineGroup)
	if err != nil {
		t.Errorf("Failed to update machine group: %v", err)
	}
}
