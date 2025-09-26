package memory

import (
	"fmt"
	"testing"
	"time"
)

func TestSet_6(t *testing.T) {
	type args struct {
		key string
		val any
		ttl time.Duration
	}
	tests := []struct {
		name string
		s    *Storage
		args args
	}{
		{
			name: "Set string value with ttl",
			s:    &Storage{data: make(map[string]item)},
			args: args{
				key: "key1",
				val: "value1",
				ttl: time.Second * 10,
			},
		},
		{
			name: "Set int value with ttl",
			s:    &Storage{data: make(map[string]item)},
			args: args{
				key: "key2",
				val: 123,
				ttl: time.Second * 10,
			},
		},
		{
			name: "Set struct value with ttl",
			s:    &Storage{data: make(map[string]item)},
			args: args{
				key: "key3",
				val: struct{ Name string }{Name: "test"},
				ttl: time.Second * 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.s.Set(tt.args.key, tt.args.val, tt.args.ttl)
			item, ok := tt.s.data[tt.args.key]
			if !ok {
				t.Errorf("Expected key %s to be in the storage", tt.args.key)
			}
			if item.v != tt.args.val {
				t.Errorf("Expected value %v, got %v", tt.args.val, item.v)
			}
			if item.e == 0 && tt.args.ttl != 0 {
				t.Errorf("Expected non-zero expiration time for key %s", tt.args.key)
			}
		})
	}
}
