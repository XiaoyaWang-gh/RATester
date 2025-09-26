package resources

import (
	"fmt"
	"testing"
)

func TestmergeData_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &genericResource{
		sd: ResourceSourceDescriptor{
			Data: make(map[string]any),
		},
	}

	in := map[string]any{
		"key1": "value1",
		"key2": "value2",
	}

	r.mergeData(in)

	if len(r.sd.Data) != 2 {
		t.Errorf("Expected 2 items in Data, got %d", len(r.sd.Data))
	}

	if r.sd.Data["key1"] != "value1" {
		t.Errorf("Expected 'value1' for key 'key1', got '%v'", r.sd.Data["key1"])
	}

	if r.sd.Data["key2"] != "value2" {
		t.Errorf("Expected 'value2' for key 'key2', got '%v'", r.sd.Data["key2"])
	}
}
