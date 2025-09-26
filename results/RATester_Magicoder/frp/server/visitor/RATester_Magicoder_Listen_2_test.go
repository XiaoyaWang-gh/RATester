package visitor

import (
	"fmt"
	"reflect"
	"testing"
)

func TestListen_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	vm := &Manager{
		listeners: make(map[string]*listenerBundle),
	}

	name := "test"
	sk := "test_sk"
	allowUsers := []string{"user1", "user2"}

	l, err := vm.Listen(name, sk, allowUsers)
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := vm.listeners[name]; !ok {
		t.Fatal("listener not added to manager")
	}

	if l == nil {
		t.Fatal("returned listener is nil")
	}

	if vm.listeners[name].sk != sk {
		t.Fatal("sk not set correctly")
	}

	if !reflect.DeepEqual(vm.listeners[name].allowUsers, allowUsers) {
		t.Fatal("allowUsers not set correctly")
	}
}
