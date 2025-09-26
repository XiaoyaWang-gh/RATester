package berror

import (
	"fmt"
	"sync"
	"testing"
)

func TestGet_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cr := &codeRegistry{
		lock:  sync.RWMutex{},
		codes: map[uint32]*codeDefinition{},
	}

	code := &codeDefinition{
		code:   1,
		module: "test",
		desc:   "test desc",
		name:   "test name",
	}

	cr.codes[code.code] = code

	testCode, ok := cr.Get(1)
	if !ok {
		t.Fatal("Expected code to be found")
	}

	if testCode.Code() != code.Code() {
		t.Fatalf("Expected code %d, got %d", code.Code(), testCode.Code())
	}

	if testCode.Module() != code.Module() {
		t.Fatalf("Expected module %s, got %s", code.Module(), testCode.Module())
	}

	if testCode.Desc() != code.Desc() {
		t.Fatalf("Expected desc %s, got %s", code.Desc(), testCode.Desc())
	}

	if testCode.Name() != code.Name() {
		t.Fatalf("Expected name %s, got %s", code.Name(), testCode.Name())
	}
}
