package alils

import (
	"fmt"
	"testing"
)

func TestDeleteMachineGroup_1(t *testing.T) {
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
	err := p.DeleteMachineGroup("test")
	if err != nil {
		t.Errorf("failed to delete machine group")
	}
}
