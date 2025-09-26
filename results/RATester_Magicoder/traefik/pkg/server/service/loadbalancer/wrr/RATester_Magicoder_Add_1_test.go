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
		fmt.Fprintln(w, "Hello, client!")
	})

	weight := 10
	b.Add("test", handler, &weight)

	if len(b.handlerMap) != 2 {
		t.Errorf("Expected 2 handlers in handlerMap, got %d", len(b.handlerMap))
	}

	if _, ok := b.handlerMap["test"]; !ok {
		t.Error("Expected 'test' handler in handlerMap")
	}

	if _, ok := b.handlerMap[hash("test")]; !ok {
		t.Error("Expected hashed 'test' handler in handlerMap")
	}

	if _, ok := b.status["test"]; !ok {
		t.Error("Expected 'test' service in status")
	}
}
