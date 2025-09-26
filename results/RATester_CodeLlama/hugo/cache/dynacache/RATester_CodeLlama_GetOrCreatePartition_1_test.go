package dynacache

import (
	"fmt"
	"testing"
)

func TestGetOrCreatePartition_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Cache{}
	p := GetOrCreatePartition[string, string](c, "foo", OptionsPartition{})
	if p == nil {
		t.Fatal("expected a partition")
	}
	if p.c == nil {
		t.Fatal("expected a cache")
	}
}
