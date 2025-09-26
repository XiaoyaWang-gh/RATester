package models

import (
	"fmt"
	"testing"
)

func TestBootstrap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mc := &ModelCache{}
	mc.Bootstrap()
}
