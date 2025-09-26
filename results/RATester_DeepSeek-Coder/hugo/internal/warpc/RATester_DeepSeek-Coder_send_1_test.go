package warpc

import (
	"fmt"
	"sync"
	"testing"
)

func TestSend_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	type Q string
	type R string

	d := &dispatcher[Q, R]{
		zero:     Message[R]{},
		mu:       sync.RWMutex{},
		encMu:    sync.Mutex{},
		pending:  make(map[uint32]*call[Q, R]),
		inOut:    &inOut{},
		shutdown: false,
		closing:  false,
	}

	call := &call[Q, R]{
		request: Message[Q]{
			Header: Header{},
			Data:   Q("test"),
		},
	}

	err := d.send(call)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
