package plugin

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestIsSupport_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &httpPlugin{
		options: v1.HTTPPluginOptions{
			Ops: []string{"op1", "op2"},
		},
	}
	if !p.IsSupport("op1") {
		t.Errorf("op1 should be supported")
	}
	if p.IsSupport("op3") {
		t.Errorf("op3 should not be supported")
	}
}
