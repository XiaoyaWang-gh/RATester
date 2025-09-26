package cache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestVacuum_1(t *testing.T) {
	type fields struct {
		RWMutex sync.RWMutex
		dur     time.Duration
		items   map[string]*MemoryItem
		Every   int
	}
	tests := []struct {
		name   string
		fields fields
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

			bc := &MemoryCache{
				RWMutex: tt.fields.RWMutex,
				dur:     tt.fields.dur,
				items:   tt.fields.items,
				Every:   tt.fields.Every,
			}
			bc.vacuum()
		})
	}
}
