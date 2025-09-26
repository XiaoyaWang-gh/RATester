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

	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	RegisterCommonFlags(fs)

	err := fs.Parse([]string{"-frpc-path=/tmp/frpc", "-frps-path=/tmp/frps", "-log-level=info", "-debug=true"})
	if err != nil {
		t.Fatalf("Failed to parse flags: %v", err)
	}

	expectedFRPClientPath := "/tmp/frpc"
	if TestContext.FRPClientPath != expectedFRPClientPath {
		t.Errorf("Expected FRPClientPath to be %q, got %q", expectedFRPClientPath, TestContext.FRPClientPath)
	}

	expectedFRPServerPath := "/tmp/frps"
	if TestContext.FRPServerPath != expectedFRPServerPath {
		t.Errorf("Expected FRPServerPath to be %q, got %q", expectedFRPServerPath, TestContext.FRPServerPath)
	}

	expectedLogLevel := "info"
	if TestContext.LogLevel != expectedLogLevel {
		t.Errorf("Expected LogLevel to be %q, got %q", expectedLogLevel, TestContext.LogLevel)
	}

	expectedDebug := true
	if TestContext.Debug != expectedDebug {
		t.Errorf("Expected Debug to be %t, got %t", expectedDebug, TestContext.Debug)
	}
}
