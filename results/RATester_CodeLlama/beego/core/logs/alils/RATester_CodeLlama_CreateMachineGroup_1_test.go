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

	p := &LogProject{
		Name:            "test",
		Endpoint:        "http://127.0.0.1:8080",
		AccessKeyID:     "accessKeyID",
		AccessKeySecret: "accessKeySecret",
	}

	m := &MachineGroup{
		Name:          "test",
		Type:          "test",
		MachineIDType: "test",
		MachineIDList: []string{"test"},
		Attribute: MachineGroupAttribute{
			ExternalName: "test",
			TopicName:    "test",
		},
	}

	err := p.CreateMachineGroup(m)
	if err != nil {
		t.Errorf("failed to create machine group")
	}
}
