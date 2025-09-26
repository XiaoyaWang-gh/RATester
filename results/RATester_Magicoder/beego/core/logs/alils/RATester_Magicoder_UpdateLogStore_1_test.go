package alils

import (
	"fmt"
	"testing"
)

func TestUpdateLogStore_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	project := &LogProject{
		Name:            "test_project",
		Endpoint:        "http://test_endpoint",
		AccessKeyID:     "test_key_id",
		AccessKeySecret: "test_key_secret",
	}

	err := project.UpdateLogStore("test_store", 30, 3)
	if err != nil {
		t.Errorf("UpdateLogStore failed: %v", err)
	}
}
