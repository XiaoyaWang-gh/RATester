package fiber

import (
	"fmt"
	"testing"
)

func TestrunOnListenHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		hooks: &Hooks{
			onListen: []func(ListenData) error{
				func(listenData ListenData) error {
					if listenData.Host != "localhost" || listenData.Port != "8080" || !listenData.TLS {
						t.Errorf("Expected listen data to be localhost:8080, got %s:%s, TLS: %t", listenData.Host, listenData.Port, listenData.TLS)
					}
					return nil
				},
			},
		},
	}

	listenData := ListenData{
		Host: "localhost",
		Port: "8080",
		TLS:  false,
	}

	app.runOnListenHooks(listenData)
}
