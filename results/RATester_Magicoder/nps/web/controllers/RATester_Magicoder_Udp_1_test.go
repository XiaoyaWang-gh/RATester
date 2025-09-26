package controllers

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestUdp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &IndexController{}
	ctrl.Udp()

	// Assert that the expected methods were called
	assert.Equal(t, "udp", ctrl.Data["info"])
	assert.Equal(t, "udp", ctrl.Data["type"])
	assert.Equal(t, "index/list", ctrl.TplName)
}
