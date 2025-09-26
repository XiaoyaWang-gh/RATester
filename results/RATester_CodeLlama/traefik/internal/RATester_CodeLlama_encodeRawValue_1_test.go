package main

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestEncodeRawValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	labels := make(map[string]string)
	root := "root"
	rawValue := map[string]interface{}{
		"key1": "value1",
		"key2": []interface{}{"value2", "value3"},
		"key3": map[string]interface{}{
			"key31": "value31",
			"key32": []interface{}{"value32", "value33"},
		},
	}

	// when
	encodeRawValue(labels, root, rawValue)

	// then
	assert.Equal(t, map[string]string{
		"root.key1":          "value1",
		"root.key2[0]":       "value2",
		"root.key2[1]":       "value3",
		"root.key3.key31":    "value31",
		"root.key3.key32[0]": "value32",
		"root.key3.key32[1]": "value33",
	}, labels)
}
