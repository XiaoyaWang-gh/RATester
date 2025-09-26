package alils

import (
	"fmt"
	"testing"
)

func TestListLogStore_1(t *testing.T) {
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

	storeNames, err := project.ListLogStore()
	if err != nil {
		t.Errorf("ListLogStore failed: %v", err)
	}

	if len(storeNames) == 0 {
		t.Errorf("ListLogStore returned no log stores")
	}
}
