package fiber

import (
	"fmt"
	"testing"
)

func TeststartupMessage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{}
	addr := "localhost:3000"
	isTLS := false
	pids := "1,2,3"
	cfg := ListenConfig{}

	app.startupMessage(addr, isTLS, pids, cfg)

	// Add assertions here
}
