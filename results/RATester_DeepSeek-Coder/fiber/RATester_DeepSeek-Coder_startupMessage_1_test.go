package fiber

import (
	"fmt"
	"testing"
)

func TestStartupMessage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	app := &App{
		config: Config{
			ColorScheme: Colors{
				Red:   "\u001b[91m",
				Reset: "\u001b[0m",
			},
		},
	}

	addr := "localhost:3000"
	isTLS := false
	pids := "1,2,3"
	cfg := ListenConfig{
		ListenerNetwork: NetworkTCP4,
	}

	app.startupMessage(addr, isTLS, pids, cfg)

	// Add your assertions here
}
