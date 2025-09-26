package memory

import (
	"fmt"
	"sync"
	"testing"
)

func TestDelete_3(t *testing.T) {
	type fields struct {
		data map[string]item
		sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "TestDelete_0",
			fields: fields{
				data: map[string]item{
					"key1": {v: "value1", e: 0},
					"key2": {v: "value2", e: 0},
				},
			},
			args: args{key: "key1"},
		},
		{
			name: "TestDelete_1",
			fields: fields{
				data: map[string]item{
					"key1": {v: "value1", e: 0},
					"key2": {v: "value2", e: 0},
				},
			},
			args: args{key: "key2"},
		},
		{
			name: "TestDelete_2",
			fields: fields{
				data: map[string]item{
					"key1": {v: "value1", e: 0},
					"key2": {v: "value2", e: 0},
				},
			},
			args: args{key: "key3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &Storage{
				data:    tt.fields.data,
				RWMutex: tt.fields.RWMutex,
			}
			s.Delete(tt.args.key)
			if _, ok := s.data[tt.args.key]; ok {
				t.Errorf("Expected key %s to be deleted", tt.args.key)
			}
		})
	}
}
