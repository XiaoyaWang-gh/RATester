package alils

import (
	"fmt"
	"testing"
)

func TestGetAppliedConfigs_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

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
		project:        &LogProject{},
	}
	confNames, err := m.GetAppliedConfigs()
	if err != nil {
		t.Errorf("GetAppliedConfigs error: %v", err)
	}
	if len(confNames) != 0 {
		t.Errorf("GetAppliedConfigs error: %v", confNames)
	}
}
