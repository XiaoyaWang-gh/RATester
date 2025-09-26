package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdd_2(t *testing.T) {
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name string
		p    PathParam
		args args
		want PathParam
	}{
		{
			name: "Add a new key-value pair",
			p:    PathParam{"existing_key": "existing_val"},
			args: args{
				key: "new_key",
				val: "new_val",
			},
			want: PathParam{
				"existing_key": "existing_val",
				"new_key":      "new_val",
			},
		},
		{
			name: "Add an existing key",
			p:    PathParam{"existing_key": "existing_val"},
			args: args{
				key: "existing_key",
				val: "new_val",
			},
			want: PathParam{
				"existing_key": "new_val",
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

			tt.p.Add(tt.args.key, tt.args.val)
			if !reflect.DeepEqual(tt.p, tt.want) {
				t.Errorf("got %v, want %v", tt.p, tt.want)
			}
		})
	}
}
