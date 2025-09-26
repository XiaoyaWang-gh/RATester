package types

import (
	"fmt"
	"testing"
)

func TestEqual_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	q1 := &BandwidthQuantity{
		s: "MB",
		i: 1024,
	}
	q2 := &BandwidthQuantity{
		s: "MB",
		i: 1024,
	}
	if !q1.Equal(q2) {
		t.Error("q1 and q2 should be equal")
	}

	q1 = &BandwidthQuantity{
		s: "KB",
		i: 1024,
	}
	q2 = &BandwidthQuantity{
		s: "MB",
		i: 1024,
	}
	if q1.Equal(q2) {
		t.Error("q1 and q2 should not be equal")
	}

	q1 = &BandwidthQuantity{
		s: "MB",
		i: 1024,
	}
	q2 = &BandwidthQuantity{
		s: "KB",
		i: 1024,
	}
	if q1.Equal(q2) {
		t.Error("q1 and q2 should not be equal")
	}

	q1 = &BandwidthQuantity{
		s: "MB",
		i: 1024,
	}
	q2 = &BandwidthQuantity{
		s: "MB",
		i: 1025,
	}
	if q1.Equal(q2) {
		t.Error("q1 and q2 should not be equal")
	}

	q1 = &BandwidthQuantity{
		s: "MB",
		i: 1024,
	}
	q2 = &BandwidthQuantity{
		s: "MB",
		i: 1024,
	}
	if !q1.Equal(q2) {
		t.Error("q1 and q2 should be equal")
	}
}
