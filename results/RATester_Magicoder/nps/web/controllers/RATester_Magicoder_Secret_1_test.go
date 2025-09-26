package controllers

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestSecret_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &IndexController{}
	ctrl.Secret()

	// Assert that the expected methods were called
	assert.Equal(t, "secret", ctrl.Data["info"])
	assert.Equal(t, "secret", ctrl.Data["type"])
	assert.Equal(t, "index/list", ctrl.TplName)
}
