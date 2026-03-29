package golang

import (
	"testing"

	"github.com/shawn-hurley/jsonrpc-golang/provider/lib"
)

func TestCommandNameUsesBinaryLocation(t *testing.T) {
	provider := NewGolangProvider(lib.Config{BinaryLocation: "/tmp/gopls"})

	if got := provider.commandName(); got != "/tmp/gopls" {
		t.Fatalf("expected binary location to be used, got %q", got)
	}
}

func TestCommandNameFallsBackToGopls(t *testing.T) {
	provider := NewGolangProvider(lib.Config{})

	if got := provider.commandName(); got != "gopls" {
		t.Fatalf("expected fallback gopls, got %q", got)
	}
}
