package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDelParams_1(t *testing.T) {
	type args struct {
		key []string
	}
	tests := []struct {
		name string
		p    PathParam
		args args
		want PathParam
	}{
		{
			name: "TestDelParams_0",
			p:    PathParam{"key1": "val1", "key2": "val2", "key3": "val3"},
			args: args{key: []string{"key1", "key3"}},
			want: PathParam{"key2": "val2"},
		},
		{
			name: "TestDelParams_1",
			p:    PathParam{"key1": "val1", "key2": "val2", "key3": "val3"},
			args: args{key: []string{"key4"}},
			want: PathParam{"key1": "val1", "key2": "val2", "key3": "val3"},
		},
		{
			name: "TestDelParams_2",
			p:    PathParam{},
			args: args{key: []string{"key1", "key2", "key3"}},
			want: PathParam{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.p.DelParams(tt.args.key...)
			if !reflect.DeepEqual(tt.p, tt.want) {
				t.Errorf("DelParams() = %v, want %v", tt.p, tt.want)
			}
		})
	}
}
