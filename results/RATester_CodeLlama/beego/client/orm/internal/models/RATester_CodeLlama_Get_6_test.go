package models

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestGet_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mc := &ModelCache{
		cache: map[string]*ModelInfo{
			"table1": {
				Name: "table1",
			},
			"table2": {
				Name: "table2",
			},
		},
	}
	mi, ok := mc.Get("table1")
	assert.True(t, ok)
	assert.Equal(t, "table1", mi.Name)
	mi, ok = mc.Get("table2")
	assert.True(t, ok)
	assert.Equal(t, "table2", mi.Name)
	mi, ok = mc.Get("table3")
	assert.False(t, ok)
	assert.Nil(t, mi)
}
