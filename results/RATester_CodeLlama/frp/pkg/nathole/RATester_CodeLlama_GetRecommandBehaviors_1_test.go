package nathole

import (
	"fmt"
	"testing"
	"time"

	"github.com/zeebo/assert"
)

func TestGetRecommandBehaviors_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := &Analyzer{
		records:             make(map[string]*MakeHoleRecords),
		dataReserveDuration: time.Second * 10,
	}

	c := &NatFeature{
		NatType:            EasyNAT,
		Behavior:           "EasyNAT",
		PortsDifference:    10,
		RegularPortsChange: true,
		PublicNetwork:      true,
	}

	v := &NatFeature{
		NatType:            EasyNAT,
		Behavior:           "EasyNAT",
		PortsDifference:    10,
		RegularPortsChange: true,
		PublicNetwork:      true,
	}

	mode, index, cBehavior, vBehavior := a.GetRecommandBehaviors("", c, v)
	assert.Equal(t, DetectMode1, mode)
	assert.Equal(t, 0, index)
	assert.Equal(t, cBehavior, vBehavior)
	assert.Equal(t, cBehavior, vBehavior)
}
