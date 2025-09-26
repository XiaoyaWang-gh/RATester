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

	project := &LogProject{
		Name:            "test_project",
		Endpoint:        "http://localhost:8080",
		AccessKeyID:     "test_key_id",
		AccessKeySecret: "test_key_secret",
	}

	err := project.DeleteMachineGroup("test_group")
	if err != nil {
		t.Errorf("DeleteMachineGroup failed: %v", err)
	}
}
