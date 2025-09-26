package main

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/cmd/frpc/sub"
	"github.com/fatedier/frp/pkg/util/system"
)

func TestMain_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	system.EnableCompatibilityMode()
	sub.Execute()
}
