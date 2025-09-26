package xlog

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestRenderPrefixString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Logger{
		prefixes: []LogPrefix{
			{
				Name:     "prefix1",
				Value:    "value1",
				Priority: 10,
			},
			{
				Name:     "prefix2",
				Value:    "value2",
				Priority: 10,
			},
			{
				Name:     "prefix3",
				Value:    "value3",
				Priority: 10,
			},
		},
	}
	l.renderPrefixString()
	assert.Equal(t, "[value1] [value2] [value3] ", l.prefixString)
}
