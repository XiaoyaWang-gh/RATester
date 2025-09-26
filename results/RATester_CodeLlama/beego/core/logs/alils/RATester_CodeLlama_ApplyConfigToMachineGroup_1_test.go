package alils

import (
	"fmt"
	"testing"
)

func TestApplyConfigToMachineGroup_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &LogProject{
		Name:            "test",
		Endpoint:        "http://127.0.0.1:8080",
		AccessKeyID:     "accessKeyID",
		AccessKeySecret: "accessKeySecret",
	}
	err := p.ApplyConfigToMachineGroup("test", "test")
	if err != nil {
		t.Errorf("ApplyConfigToMachineGroup failed, err: %v", err)
	}
}
