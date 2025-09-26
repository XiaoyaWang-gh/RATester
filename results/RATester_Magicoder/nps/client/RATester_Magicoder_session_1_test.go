package client

import (
	"container/heap"
	"fmt"
	"testing"
	"time"

	"ehang.io/nps/lib/file"
	"ehang.io/nps/lib/sheap"
)

func Testsession_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	healths := []*file.Health{
		{
			HealthCheckTimeout:  10,
			HealthMaxFail:       3,
			HealthCheckInterval: 5,
			HealthNextTime:      time.Now().Add(time.Duration(5) * time.Second),
			HealthMap:           make(map[string]int),
			HttpHealthUrl:       "http://localhost:8080",
			HealthRemoveArr:     []string{},
			HealthCheckType:     "http",
			HealthCheckTarget:   "localhost:8080",
		},
	}

	h := &sheap.IntHeap{}
	heap.Init(h)
	heap.Push(h, time.Now().Add(time.Duration(5)*time.Second).Unix())

	session(healths, h)

	if h.Len() != 0 {
		t.Errorf("Expected 0, got %d", h.Len())
	}
}
