package fiber

import (
	"fmt"
	"testing"
)

func TestRunOnListenHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		hooks: &Hooks{
			onListen: []func(ListenData) error{
				func(data ListenData) error {
					if data.Host != "localhost" || data.Port != "3000" || !data.TLS {
						t.Errorf("Expected data to be localhost:3000, TLS: true, got %s:%s, TLS: %v", data.Host, data.Port, data.TLS)
					}
					return nil
				},
			},
		},
	}

	listenData := ListenData{
		Host: "localhost",
		Port: "3000",
		TLS:  true,
	}

	app.runOnListenHooks(listenData)
}
