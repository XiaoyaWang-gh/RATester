package validation

import (
	"fmt"
	"testing"
)

func TestBase64_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &Validation{}
	obj := "1234567890"
	key := "key"
	r := v.Base64(obj, key)
	if r.Error == nil {
		t.Error("Expected error")
	}
	if r.Error.Key != key {
		t.Errorf("Expected key %s, got %s", key, r.Error.Key)
	}
	if r.Error.Message != "key must be base64 encoded" {
		t.Errorf("Expected message %s, got %s", "key must be base64 encoded", r.Error.Message)
	}
}
