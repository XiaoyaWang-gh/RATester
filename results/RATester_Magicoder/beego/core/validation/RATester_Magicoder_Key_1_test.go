package validation

import (
	"fmt"
	"testing"
)

func TestKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Result{
		Error: &Error{
			Key:     "testKey",
			Field:   "testField",
			Message: "testMessage",
		},
		Ok: true,
	}

	key := "newKey"
	r.Key(key)

	if r.Error.Key != key {
		t.Errorf("Expected Key to be %s, but got %s", key, r.Error.Key)
	}
}
