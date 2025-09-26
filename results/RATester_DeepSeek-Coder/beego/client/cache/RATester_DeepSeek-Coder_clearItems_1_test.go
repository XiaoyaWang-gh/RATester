package cache

import (
	"fmt"
	"testing"
)

func TestClearItems_1(t *testing.T) {
	type args struct {
		keys []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test clearItems with one key",
			args: args{keys: []string{"key1"}},
		},
		{
			name: "Test clearItems with multiple keys",
			args: args{keys: []string{"key1", "key2", "key3"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			bc := &MemoryCache{
				items: map[string]*MemoryItem{
					"key1": {val: "value1"},
					"key2": {val: "value2"},
					"key3": {val: "value3"},
				},
			}
			bc.clearItems(tt.args.keys)
			for _, key := range tt.args.keys {
				if _, ok := bc.items[key]; ok {
					t.Errorf("Expected key %s to be deleted", key)
				}
			}
		})
	}
}
