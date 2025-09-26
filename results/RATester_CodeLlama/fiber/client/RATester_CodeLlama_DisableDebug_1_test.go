package client

import (
	"fmt"
	"testing"
)

func TestDisableDebug_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Client{}
	c.DisableDebug()
	if c.debug != false {
		t.Errorf("DisableDebug() failed")
	}
}
