package client

import (
	"fmt"
	"testing"
)

func TestSetKeyValueBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cj := &CookieJar{}
	host := "www.example.com"
	key := []byte("key")
	value := []byte("value")
	cj.SetKeyValueBytes(host, key, value)
}
