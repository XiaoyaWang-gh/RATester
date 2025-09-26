package config

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewFakeConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fakeConfig := NewFakeConfig()
	assert.NotNil(t, fakeConfig)
}
