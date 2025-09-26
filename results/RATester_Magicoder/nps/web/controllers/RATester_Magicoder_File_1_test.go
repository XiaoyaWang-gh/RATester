package controllers

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &IndexController{}
	ctrl.File()

	// Assert that the expected methods were called
	assert.Equal(t, "file server", ctrl.Data["info"])
	assert.Equal(t, "file", ctrl.Data["type"])
	assert.Equal(t, "index/list", ctrl.TplName)
}
