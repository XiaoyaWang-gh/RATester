package web

import (
	"fmt"
	"testing"
)

func TestregisterHttpServer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &adminController{}
	svr := &HttpServer{}
	ctrl.registerHttpServer(svr)

	if len(ctrl.servers) != 1 {
		t.Errorf("Expected length of servers to be 1, but got %d", len(ctrl.servers))
	}

	if ctrl.servers[0] != svr {
		t.Errorf("Expected server to be %v, but got %v", svr, ctrl.servers[0])
	}
}
