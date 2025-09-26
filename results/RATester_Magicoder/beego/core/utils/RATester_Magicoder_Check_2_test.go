package utils

import (
	"fmt"
	"sync"
	"testing"
)

func TestCheck_2(t *testing.T) {
	bm := &BeeMap{
		lock: new(sync.RWMutex),
		bm:   make(map[interface{}]interface{}),
	}

	bm.Set("key1", "value1")

	tests := []struct {
		name string
		key  interface{}
		want bool
	}{
		{"Test Case 1", "key1", true},
		{"Test Case 2", "key2", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := bm.Check(tt.key); got != tt.want {
				t.Errorf("BeeMap.Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
