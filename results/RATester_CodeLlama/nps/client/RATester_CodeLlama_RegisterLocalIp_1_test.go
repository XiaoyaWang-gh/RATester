package client

import (
	"fmt"
	"testing"
)

func TestRegisterLocalIp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	RegisterLocalIp("127.0.0.1:8080", "1234567890", "tcp", "http://127.0.0.1:8080", 1)
}
