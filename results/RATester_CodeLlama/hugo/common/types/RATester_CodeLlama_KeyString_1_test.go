package types

import (
	"fmt"
	"testing"
)

func TestKeyString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	k := KeyValues{Key: "key"}
	if k.KeyString() != "key" {
		t.Errorf("KeyString() = %q, want %q", k.KeyString(), "key")
	}
}
