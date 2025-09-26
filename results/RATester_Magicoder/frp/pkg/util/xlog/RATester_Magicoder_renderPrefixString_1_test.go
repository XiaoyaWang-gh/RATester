package xlog

import (
	"fmt"
	"testing"
)

func TestRenderPrefixString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Logger{
		prefixes: []LogPrefix{
			{Name: "prefix1", Value: "value1", Priority: 1},
			{Name: "prefix2", Value: "value2", Priority: 2},
			{Name: "prefix3", Value: "value3", Priority: 3},
		},
	}
	l.renderPrefixString()
	if l.prefixString != "[value1] [value2] [value3] " {
		t.Errorf("Expected prefixString to be '[value1] [value2] [value3] ', but got '%s'", l.prefixString)
	}
}
