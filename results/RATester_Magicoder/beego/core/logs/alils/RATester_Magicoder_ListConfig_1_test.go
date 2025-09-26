package alils

import (
	"fmt"
	"testing"
)

func TestListConfig_1(t *testing.T) {
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

	cfgNames, total, err := project.ListConfig(0, 100)
	if err != nil {
		t.Errorf("ListConfig failed: %v", err)
	}

	if total < 0 {
		t.Errorf("ListConfig failed: total should be greater than or equal to 0, but got %v", total)
	}

	if len(cfgNames) != total {
		t.Errorf("ListConfig failed: length of cfgNames should be equal to total, but got %v and %v", len(cfgNames), total)
	}
}
