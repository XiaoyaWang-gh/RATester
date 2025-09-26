package visitor

import (
	"fmt"
	"testing"
)

func TestListen_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	vm := &Manager{
		listeners: make(map[string]*listenerBundle),
	}

	name := "test"
	sk := "test_sk"
	allowUsers := []string{"user1", "user2"}

	l, err := vm.Listen(name, sk, allowUsers)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, ok := vm.listeners[name]; !ok {
		t.Fatalf("listener not found in manager")
	}

	if l == nil {
		t.Fatalf("expected listener, got nil")
	}

	_, err = vm.Listen(name, sk, allowUsers)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	expectedErr := fmt.Errorf("custom listener for [%s] is repeated", name)
	if err.Error() != expectedErr.Error() {
		t.Fatalf("expected error %v, got %v", expectedErr, err)
	}
}
