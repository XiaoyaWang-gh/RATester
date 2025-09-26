package test

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestSys_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	fi := bindataFileInfo{
		name: "test",
	}

	// when
	result := fi.Sys()

	// then
	assert.Nil(t, result)
}
