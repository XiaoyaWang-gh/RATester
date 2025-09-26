package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetByFullName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mc := &ModelCache{}
	mi := &ModelInfo{}
	mc.cacheByFullName = map[string]*ModelInfo{"name": mi}
	mi, ok := mc.GetByFullName("name")
	assert.Equal(t, mi, mi)
	assert.True(t, ok)
}
