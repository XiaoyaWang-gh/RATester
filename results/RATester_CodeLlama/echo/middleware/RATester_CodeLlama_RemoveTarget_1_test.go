package middleware

import (
	"fmt"
	"testing"
)

func TestRemoveTarget_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &commonBalancer{}
	b.AddTarget(&ProxyTarget{Name: "target1"})
	b.AddTarget(&ProxyTarget{Name: "target2"})
	b.AddTarget(&ProxyTarget{Name: "target3"})
	if !b.RemoveTarget("target2") {
		t.Error("RemoveTarget failed")
	}
	if b.RemoveTarget("target2") {
		t.Error("RemoveTarget failed")
	}
	if b.RemoveTarget("target4") {
		t.Error("RemoveTarget failed")
	}
}
