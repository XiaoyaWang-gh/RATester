package maps

import (
	"fmt"
	"sync"
	"testing"
)

func TestDelete_3(t *testing.T) {
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
		want   string
	}{
		{
			name: "TestDelete_0",
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
			want: "",
		},
		{
			name: "TestDelete_1",
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
			want: "",
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
			if got := c.Delete(tt.args.key); got != tt.want {
				t.Errorf("Scratch.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
