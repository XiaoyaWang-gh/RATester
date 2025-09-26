package main

import (
	"fmt"
	"testing"
)

func TestLoad_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := client{}
	parentKey := "parentKey"
	conf := map[string]interface{}{
		"key1": map[string]interface{}{
			"key11": "value11",
			"key12": "value12",
		},
		"key2": []map[string]interface{}{
			{
				"key21": "value21",
				"key22": "value22",
			},
			{
				"key23": "value23",
				"key24": "value24",
			},
		},
		"key3": []interface{}{
			"value31",
			"value32",
		},
		"key4": "value4",
	}

	// when
	err := c.load(parentKey, conf)

	// then
	if err != nil {
		t.Errorf("load() should not return an error, but it does: %v", err)
	}
}
