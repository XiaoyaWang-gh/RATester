package cache

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/internal/memory"
)

func TestAcquire_2(t *testing.T) {
	type fields struct {
		pool    sync.Pool
		memory  *memory.Storage
		storage fiber.Storage
	}
	tests := []struct {
		name   string
		fields fields
		want   *item
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := &manager{
				pool:    tt.fields.pool,
				memory:  tt.fields.memory,
				storage: tt.fields.storage,
			}
			if got := m.acquire(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("manager.acquire() = %v, want %v", got, tt.want)
			}
		})
	}
}
