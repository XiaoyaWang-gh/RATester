package gin

import (
	"fmt"
	"net"
	"net/http"
	"testing"
)

func TestRunListener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := New()
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		err = engine.RunListener(listener)
		if err != nil {
			t.Error(err)
		}
	}()

	resp, err := http.Get(fmt.Sprintf("http://localhost:%d", listener.Addr().(*net.TCPAddr).Port))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
