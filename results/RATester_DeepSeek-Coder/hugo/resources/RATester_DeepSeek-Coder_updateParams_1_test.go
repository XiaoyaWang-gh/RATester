package resources

import (
	"fmt"
	"testing"
)

func TestUpdateParams_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &metaResource{
		changed: false,
		title:   "test",
		name:    "test",
		params:  make(map[string]interface{}),
	}

	params := map[string]any{
		"key1": "value1",
		"key2": 2,
		"key3": true,
	}

	r.updateParams(params)

	if !r.changed {
		t.Errorf("Expected changed to be true after updateParams, but it was false")
	}

	if len(r.params) != len(params) {
		t.Errorf("Expected params to have %d elements, but it has %d", len(params), len(r.params))
	}

	for k, v := range params {
		if r.params[k] != v {
			t.Errorf("Expected params[%s] to be %v, but it was %v", k, v, r.params[k])
		}
	}
}
