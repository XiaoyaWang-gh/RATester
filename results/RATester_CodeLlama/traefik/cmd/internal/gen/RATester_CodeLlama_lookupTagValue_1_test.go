package main

import (
	"fmt"
	"testing"
)

func TestLookupTagValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	raw := "json:\"name,omitempty\""
	key := "json"
	values, ok := lookupTagValue(raw, key)
	if !ok {
		t.Errorf("lookupTagValue(%q, %q) = %v, %v; want %v, %v", raw, key, values, ok, []string{"name"}, true)
	}
}
