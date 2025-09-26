package cache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Testrelease_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := &manager{
		pool: sync.Pool{
			New: func() interface{} {
				return &item{}
			},
		},
	}

	item := manager.acquire()
	item.body = []byte("test")
	item.ctype = []byte("text/plain")
	item.status = 200
	item.exp = uint64(time.Now().Unix())
	item.headers = map[string][]byte{"header1": []byte("value1")}

	manager.release(item)

	if item.body != nil || item.ctype != nil || item.status != 0 || item.exp != 0 || item.headers != nil {
		t.Errorf("Expected item to be reset after release, but it's not")
	}
}
