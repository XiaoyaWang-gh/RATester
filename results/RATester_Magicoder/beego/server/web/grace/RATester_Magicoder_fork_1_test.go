package grace

import (
	"fmt"
	"net"
	"net/http"
	"testing"
)

func Testfork_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	srv := &Server{
		Server: &http.Server{
			Addr: ":8080",
		},
		ln: &net.TCPListener{},
	}

	err := srv.fork()
	if err != nil {
		t.Errorf("fork() returned an error: %v", err)
	}
}
