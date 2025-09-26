package hexec

import (
	"fmt"
	"sync"
	"testing"
)

func TestNpx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Exec{}
	e.checkNpx()
	if e.npxAvailable {
		t.Error("npxAvailable should be false")
	}
	e.npxInit = sync.Once{}
	e.npxAvailable = true
	if _, err := e.npx(""); err != nil {
		t.Error(err)
	}
}
