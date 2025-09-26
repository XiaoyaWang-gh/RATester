package wrr

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAdd_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &Balancer{
		handlerMap: make(map[string]*namedHandler),
		status:     make(map[string]struct{}),
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	})

	weight := 5
	b.Add("testHandler", handler, &weight)

	b.handlersMu.RLock()
	defer b.handlersMu.RUnlock()

	if len(b.handlerMap) != 2 {
		t.Errorf("Expected 2 handlers in handlerMap, got %d", len(b.handlerMap))
	}

	if _, ok := b.handlerMap["testHandler"]; !ok {
		t.Error("Expected handlerMap to contain 'testHandler'")
	}

	if _, ok := b.handlerMap[hash("testHandler")]; !ok {
		t.Error("Expected handlerMap to contain hashed value of 'testHandler'")
	}

	if _, ok := b.status["testHandler"]; !ok {
		t.Error("Expected status to contain 'testHandler'")
	}
}
