package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewKVs_1(t *testing.T) {
	type args struct {
		kvs []KV
	}
	tests := []struct {
		name string
		args args
		want KVs
	}{
		{
			name: "Test case 1",
			args: args{
				kvs: []KV{
					&SimpleKV{
						Key:   "key1",
						Value: "value1",
					},
					&SimpleKV{
						Key:   "key2",
						Value: "value2",
					},
				},
			},
			want: &SimpleKVs{
				kvs: map[interface{}]interface{}{
					"key1": "value1",
					"key2": "value2",
				},
			},
		},
		{
			name: "Test case 2",
			args: args{
				kvs: []KV{
					&SimpleKV{
						Key:   "key3",
						Value: "value3",
					},
					&SimpleKV{
						Key:   "key4",
						Value: "value4",
					},
				},
			},
			want: &SimpleKVs{
				kvs: map[interface{}]interface{}{
					"key3": "value3",
					"key4": "value4",
				},
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

			if got := NewKVs(tt.args.kvs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKVs() = %v, want %v", got, tt.want)
			}
		})
	}
}
