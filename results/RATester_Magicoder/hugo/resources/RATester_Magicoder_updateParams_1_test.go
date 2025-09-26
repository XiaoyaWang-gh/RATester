package resources

import (
	"fmt"
	"testing"
)

func TestupdateParams_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &metaResource{
		changed: false,
		title:   "Test Resource",
		name:    "test",
		params:  make(map[string]any),
	}

	params := map[string]any{
		"key1": "value1",
		"key2": 2,
	}

	r.updateParams(params)

	if !r.changed {
		t.Error("Expected changed to be true after updateParams")
	}

	if r.params["key1"] != "value1" {
		t.Errorf("Expected key1 to be value1, got %v", r.params["key1"])
	}

	if r.params["key2"] != 2 {
		t.Errorf("Expected key2 to be 2, got %v", r.params["key2"])
	}
}
