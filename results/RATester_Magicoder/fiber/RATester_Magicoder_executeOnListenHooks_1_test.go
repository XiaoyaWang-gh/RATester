package fiber

import (
	"fmt"
	"testing"
)

func TestexecuteOnListenHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hooks := &Hooks{}

	listenData := ListenData{
		Host: "localhost",
		Port: "3000",
		TLS:  false,
	}

	hooks.OnListen(func(data ListenData) error {
		if data.Host != listenData.Host {
			t.Errorf("Expected host %s, got %s", listenData.Host, data.Host)
		}
		if data.Port != listenData.Port {
			t.Errorf("Expected port %s, got %s", listenData.Port, data.Port)
		}
		if data.TLS != listenData.TLS {
			t.Errorf("Expected TLS %t, got %t", listenData.TLS, data.TLS)
		}
		return nil
	})

	err := hooks.executeOnListenHooks(listenData)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
