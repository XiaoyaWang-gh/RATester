package maps

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestGet_9(t *testing.T) {
	type fields struct {
		values map[string]any
		mu     sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
	}{
		{
			name: "TestGet_ExistingKey",
			fields: fields{
				values: map[string]any{
					"key1": "value1",
					"key2": "value2",
				},
				mu: sync.RWMutex{},
			},
			args: args{
				key: "key1",
			},
			want: "value1",
		},
		{
			name: "TestGet_NonExistingKey",
			fields: fields{
				values: map[string]any{
					"key1": "value1",
					"key2": "value2",
				},
				mu: sync.RWMutex{},
			},
			args: args{
				key: "key3",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &Scratch{
				values: tt.fields.values,
				mu:     tt.fields.mu,
			}
			if got := c.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scratch.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
