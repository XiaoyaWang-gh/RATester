package alils

import (
	"fmt"
	"testing"
)

func TestRemoveConfigFromMachineGroup_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &LogProject{
		Name:            "test",
		Endpoint:        "127.0.0.1:8080",
		AccessKeyID:     "access_key_id",
		AccessKeySecret: "access_key_secret",
	}
	confName := "test_config"
	groupName := "test_group"
	err := p.RemoveConfigFromMachineGroup(confName, groupName)
	if err != nil {
		t.Errorf("RemoveConfigFromMachineGroup failed, err: %v", err)
	}
}
