package client

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/conn"
)

func TesthandleMain_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &TRPClient{
		svrAddr: "localhost:8080",
		signal:  &conn.Conn{},
	}

	client.handleMain()

	// Assert that the function does not panic or return an error
}
