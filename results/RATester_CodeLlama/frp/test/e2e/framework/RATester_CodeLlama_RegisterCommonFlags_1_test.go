package framework

import (
	"flag"
	"fmt"
	"testing"
)

func TestRegisterCommonFlags_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	RegisterCommonFlags(flag.NewFlagSet("test", flag.ContinueOnError))
	if TestContext.FRPClientPath != "../../bin/frpc" {
		t.Errorf("TestContext.FRPClientPath should be '../../bin/frpc', but got %s", TestContext.FRPClientPath)
	}
	if TestContext.FRPServerPath != "../../bin/frps" {
		t.Errorf("TestContext.FRPServerPath should be '../../bin/frps', but got %s", TestContext.FRPServerPath)
	}
	if TestContext.LogLevel != "debug" {
		t.Errorf("TestContext.LogLevel should be 'debug', but got %s", TestContext.LogLevel)
	}
	if TestContext.Debug != false {
		t.Errorf("TestContext.Debug should be 'false', but got %v", TestContext.Debug)
	}
}
