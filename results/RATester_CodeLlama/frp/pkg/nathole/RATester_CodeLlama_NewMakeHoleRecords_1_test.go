package nathole

import (
	"fmt"
	"testing"
)

func TestNewMakeHoleRecords_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &NatFeature{
		NatType:            "1",
		Behavior:           "1",
		PortsDifference:    1,
		RegularPortsChange: true,
		PublicNetwork:      true,
	}
	v := &NatFeature{
		NatType:            "2",
		Behavior:           "2",
		PortsDifference:    2,
		RegularPortsChange: false,
		PublicNetwork:      false,
	}
	mhr := NewMakeHoleRecords(c, v)
	if len(mhr.scores) != 3 {
		t.Errorf("NewMakeHoleRecords() = %v, want %v", len(mhr.scores), 3)
	}
}
