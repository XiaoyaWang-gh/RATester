package nathole

import (
	"fmt"
	"testing"
)

func TestGetRecommandBehaviors_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := &Analyzer{
		records: make(map[string]*MakeHoleRecords),
	}

	key := "192.168.1.1:192.168.1.2"
	c := &NatFeature{
		NatType:            "EasyNAT",
		Behavior:           "Behavior1",
		PortsDifference:    10,
		RegularPortsChange: true,
		PublicNetwork:      true,
	}
	v := &NatFeature{
		NatType:            "HardNAT",
		Behavior:           "Behavior2",
		PortsDifference:    20,
		RegularPortsChange: false,
		PublicNetwork:      false,
	}

	mode, index, cBehavior, vBehavior := a.GetRecommandBehaviors(key, c, v)

	if mode != 1 || index != 1 {
		t.Errorf("Expected mode and index to be 1, 1 but got %d, %d", mode, index)
	}

	if cBehavior.Role != "Role1" || vBehavior.Role != "Role2" {
		t.Errorf("Expected cBehavior and vBehavior to have Role1 and Role2 but got %s, %s", cBehavior.Role, vBehavior.Role)
	}
}
