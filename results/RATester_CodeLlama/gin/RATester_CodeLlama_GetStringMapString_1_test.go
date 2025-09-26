package gin

import (
	"fmt"
	"testing"
)

func TestGetStringMapString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Keys: map[string]any{
			"key": map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
	}

	sms := c.GetStringMapString("key")
	if sms == nil {
		t.Errorf("GetStringMapString() = %v, want %v", sms, nil)
	}

	if sms["key1"] != "value1" {
		t.Errorf("GetStringMapString() = %v, want %v", sms["key1"], "value1")
	}

	if sms["key2"] != "value2" {
		t.Errorf("GetStringMapString() = %v, want %v", sms["key2"], "value2")
	}
}
