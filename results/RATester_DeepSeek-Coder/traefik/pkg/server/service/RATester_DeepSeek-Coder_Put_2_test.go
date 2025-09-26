package service

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPut_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &bufferPool{}
	bytes := []byte("test")
	b.Put(bytes)

	got := b.pool.Get()
	if !reflect.DeepEqual(got, bytes) {
		t.Errorf("Expected %v, got %v", bytes, got)
	}
}
