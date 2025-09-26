package controllers

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestTcp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &IndexController{}
	ctrl.Tcp()

	// Assert that the expected methods were called
	assert.Equal(t, "tcp", ctrl.Data["info"])
	assert.Equal(t, "tcp", ctrl.Data["type"])
	assert.Equal(t, "index/list", ctrl.TplName)
}
