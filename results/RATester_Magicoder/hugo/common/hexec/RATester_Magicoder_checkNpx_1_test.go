package hexec

import (
	"fmt"
	"sync"
	"testing"
)

func TestcheckNpx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Exec{
		npxInit: sync.Once{},
	}

	e.checkNpx()

	if !e.npxAvailable {
		t.Errorf("Expected npx to be available, but it was not")
	}
}
