package grace

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestListenAndServe_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	addr := "localhost:8080"
	go func() {
		err := ListenAndServe(addr, handler)
		if err != nil {
			t.Errorf("ListenAndServe returned an error: %v", err)
		}
	}()

	time.Sleep(100 * time.Millisecond)

	resp, err := http.Get("http://" + addr)
	if err != nil {
		t.Errorf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Failed to read response body: %v", err)
	}

	if string(body) != "Hello, world!" {
		t.Errorf("Expected 'Hello, world!', got '%s'", string(body))
	}
}
