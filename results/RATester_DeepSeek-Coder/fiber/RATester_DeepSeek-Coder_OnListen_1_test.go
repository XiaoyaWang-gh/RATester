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
	hooks := &Hooks{app: app}

	handler := func(ld ListenData) error {
		if ld.Host != "localhost" || ld.Port != "3000" || ld.TLS != true {
			t.Errorf("Expected ListenData to be localhost:3000:true, got %s:%s:%t", ld.Host, ld.Port, ld.TLS)
		}
		return nil
	}

	hooks.OnListen(handler)

	hooks.executeOnListenHooks(ListenData{Host: "localhost", Port: "3000", TLS: true})
}
