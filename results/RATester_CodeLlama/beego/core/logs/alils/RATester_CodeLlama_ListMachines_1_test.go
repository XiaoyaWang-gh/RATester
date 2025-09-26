package alils

import (
	"fmt"
	"testing"
)

func TestListMachines_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &MachineGroup{
		Name:          "test",
		Type:          "host",
		MachineIDType: "ip",
		MachineIDList: []string{"127.0.0.1"},
		Attribute: MachineGroupAttribute{
			ExternalName: "test",
			TopicName:    "test",
		},
	}

	ms, total, err := m.ListMachines()
	if err != nil {
		t.Errorf("ListMachines failed, err: %v", err)
	}

	if total != 1 {
		t.Errorf("ListMachines failed, total: %v", total)
	}

	if len(ms) != 1 {
		t.Errorf("ListMachines failed, ms: %v", ms)
	}

	if ms[0].IP != "127.0.0.1" {
		t.Errorf("ListMachines failed, ms[0].IP: %v", ms[0].IP)
	}

	if ms[0].UniqueID != "127.0.0.1" {
		t.Errorf("ListMachines failed, ms[0].UniqueID: %v", ms[0].UniqueID)
	}

	if ms[0].UserdefinedID != "" {
		t.Errorf("ListMachines failed, ms[0].UserdefinedID: %v", ms[0].UserdefinedID)
	}
}
