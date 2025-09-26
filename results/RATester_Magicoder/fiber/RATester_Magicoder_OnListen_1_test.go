package fiber

import (
	"fmt"
	"testing"
)

func TestOnListen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	hooks := app.Hooks()

	handler := func(listenData ListenData) error {
		if listenData.Host != "localhost" {
			t.Errorf("Expected host to be 'localhost', got %s", listenData.Host)
		}
		if listenData.Port != "3000" {
			t.Errorf("Expected port to be '3000', got %s", listenData.Port)
		}
		if listenData.TLS != false {
			t.Errorf("Expected TLS to be false, got %t", listenData.TLS)
		}
		return nil
	}

	hooks.OnListen(handler)

	listenData := ListenData{
		Host: "localhost",
		Port: "3000",
		TLS:  false,
	}

	hooks.executeOnListenHooks(listenData)
}
