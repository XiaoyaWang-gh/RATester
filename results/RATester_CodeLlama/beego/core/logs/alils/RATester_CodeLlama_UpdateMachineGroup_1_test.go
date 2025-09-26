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

	p := &LogProject{
		Name:            "test",
		Endpoint:        "127.0.0.1:8080",
		AccessKeyID:     "access_key_id",
		AccessKeySecret: "access_key_secret",
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
		CreateTime:     1,
		LastModifyTime: 1,
		project:        p,
	}

	err := p.UpdateMachineGroup(m)
	if err != nil {
		t.Errorf("failed to update machine group")
	}
}
