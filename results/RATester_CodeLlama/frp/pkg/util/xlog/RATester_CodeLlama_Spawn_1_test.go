package xlog

import (
	"fmt"
	"testing"
)

func TestSpawn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := New()
	l.AppendPrefix("prefix1")
	l.AppendPrefix("prefix2")
	l.renderPrefixString()
	nl := l.Spawn()
	nl.AppendPrefix("prefix3")
	nl.renderPrefixString()
	if nl.prefixString != "prefix1prefix2prefix3" {
		t.Errorf("nl.prefixString = %s, want prefix1prefix2prefix3", nl.prefixString)
	}
}
