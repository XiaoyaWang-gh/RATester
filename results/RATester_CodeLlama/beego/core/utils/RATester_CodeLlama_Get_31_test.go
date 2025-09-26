package utils

import (
	"fmt"
	"sync"
	"testing"
)

func TestGet_31(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &BeeMap{
		lock: &sync.RWMutex{},
		bm:   map[interface{}]interface{}{"k": "v"},
	}
	if m.Get("k") != "v" {
		t.Errorf("Get failed")
	}
}
