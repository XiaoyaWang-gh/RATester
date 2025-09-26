package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given:
	mc := &ModelCache{}
	table := "table"
	mi := &ModelInfo{}

	// when:
	mc.Set(table, mi)

	// then:
	assert.Equal(t, mi, mc.cache[table])
	assert.Equal(t, mi, mc.cacheByFullName[mi.FullName])
	assert.Equal(t, []string{table}, mc.orders)
}
