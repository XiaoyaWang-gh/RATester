package logger

import (
	"fmt"
	"testing"
)

func TestbeforeHandlerFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cfg := Config{
		enableColors: true,
	}

	beforeHandlerFunc(cfg)

	if cfg.Output == nil {
		t.Error("Output should not be nil")
	}
}
