package common

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPut_1(t *testing.T) {
	pool := &copyBufferPool{}

	t.Run("puts_correct_size_buffer_to_pool", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		buffer := make([]byte, PoolSizeCopy)
		pool.Put(buffer)
		got := pool.pool.Get()
		if !reflect.DeepEqual(got, buffer) {
			t.Errorf("Expected %v, got %v", buffer, got)
		}
	})

	t.Run("puts_incorrect_size_buffer_to_nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		buffer := make([]byte, PoolSizeCopy+1)
		pool.Put(buffer)
		if pool.pool.Get() != nil {
			t.Errorf("Expected nil, got %v", pool.pool.Get())
		}
	})
}
