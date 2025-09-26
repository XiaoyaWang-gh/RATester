package fiber

import (
	"errors"
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestExecuteOnListenHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	h := &Hooks{}
	h.onListen = []OnListenHandler{
		func(listenData ListenData) error {
			if listenData.Host != "localhost" {
				return errors.New("host is not localhost")
			}
			return nil
		},
		func(listenData ListenData) error {
			if listenData.Port != "8080" {
				return errors.New("port is not 8080")
			}
			return nil
		},
	}
	listenData := ListenData{
		Host: "localhost",
		Port: "8080",
	}

	// when
	err := h.executeOnListenHooks(listenData)

	// then
	assert.NoError(t, err)
}
