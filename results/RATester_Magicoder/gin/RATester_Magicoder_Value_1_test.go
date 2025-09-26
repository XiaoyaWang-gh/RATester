package gin

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestValue_1(t *testing.T) {
	type args struct {
		key any
	}
	tests := []struct {
		name string
		c    *Context
		args args
		want any
	}{
		{
			name: "Test case 1",
			c: &Context{
				Request: &http.Request{},
				Keys: map[string]any{
					"key1": "value1",
				},
			},
			args: args{key: "key1"},
			want: "value1",
		},
		{
			name: "Test case 2",
			c: &Context{
				Request: &http.Request{},
				Keys: map[string]any{
					"key2": "value2",
				},
			},
			args: args{key: "key2"},
			want: "value2",
		},
		{
			name: "Test case 3",
			c: &Context{
				Request: &http.Request{},
				Keys: map[string]any{
					"key3": "value3",
				},
			},
			args: args{key: "key3"},
			want: "value3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.Value(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Context.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
