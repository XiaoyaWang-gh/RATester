package middleware

import (
	"fmt"
	"testing"
)

func TestAddTarget_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &commonBalancer{}
	target := &ProxyTarget{Name: "target1"}
	if !b.AddTarget(target) {
		t.Errorf("AddTarget() = %v, want %v", false, true)
	}
	if b.targets[0].Name != "target1" {
		t.Errorf("AddTarget() = %v, want %v", b.targets[0].Name, "target1")
	}
	if b.AddTarget(target) {
		t.Errorf("AddTarget() = %v, want %v", true, false)
	}
}
