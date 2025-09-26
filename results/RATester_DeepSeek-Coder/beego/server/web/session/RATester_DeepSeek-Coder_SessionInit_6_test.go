package session

import (
	"container/list"
	"context"
	"fmt"
	"sync"
	"testing"
)

func TestSessionInit_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	pder := &MemProvider{
		lock:        sync.RWMutex{},
		sessions:    make(map[string]*list.Element),
		list:        list.New(),
		maxlifetime: 0,
		savePath:    "",
	}

	maxlifetime := int64(3600)
	savePath := "/tmp"

	err := pder.SessionInit(ctx, maxlifetime, savePath)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if pder.maxlifetime != maxlifetime {
		t.Errorf("Expected %v, got %v", maxlifetime, pder.maxlifetime)
	}

	if pder.savePath != savePath {
		t.Errorf("Expected %v, got %v", savePath, pder.savePath)
	}
}
