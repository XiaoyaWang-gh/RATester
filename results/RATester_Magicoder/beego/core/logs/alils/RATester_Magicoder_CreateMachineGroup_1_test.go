package alils

import (
	"fmt"
	"testing"
)

func TestCreateMachineGroup_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	project := &LogProject{
		Name:            "test_project",
		Endpoint:        "http://localhost:8080",
		AccessKeyID:     "test_id",
		AccessKeySecret: "test_secret",
	}

	machineGroup := &MachineGroup{
		Name:          "test_group",
		Type:          "test_type",
		MachineIDType: "test_id_type",
		MachineIDList: []string{"test_id1", "test_id2"},
		Attribute: MachineGroupAttribute{
			ExternalName: "test_external_name",
			TopicName:    "test_topic_name",
		},
		project: project,
	}

	err := project.CreateMachineGroup(machineGroup)
	if err != nil {
		t.Errorf("Failed to create machine group: %v", err)
	}
}
