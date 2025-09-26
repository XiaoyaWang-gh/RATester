package service

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestPut_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pool := &bufferPool{
		pool: sync.Pool{
			New: func() interface{} {
				return make([]byte, 1024)
			},
		},
	}

	bytes := pool.Get()
	pool.Put(bytes)

	got := pool.pool.Get()
	if !reflect.DeepEqual(bytes, got) {
		t.Errorf("Put() = %v, want %v", got, bytes)
	}
}
