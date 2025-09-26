package fiber

import (
	"fmt"
	"testing"
)

func TestConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		config: Config{
			ServerHeader: "TestServer",
		},
	}

	cfg := app.Config()

	if cfg.ServerHeader != "TestServer" {
		t.Errorf("Expected ServerHeader to be 'TestServer', but got '%s'", cfg.ServerHeader)
	}
}
