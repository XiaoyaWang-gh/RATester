package alils

import (
	"fmt"
	"testing"
)

func TestListMachineGroup_1(t *testing.T) {
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

	_, _, err := project.ListMachineGroup(0, 10)
	if err != nil {
		t.Errorf("ListMachineGroup failed, err: %v", err)
	}
}
