package gin

import (
	"fmt"
	"testing"
)

func TestSet_1(t *testing.T) {
	type args struct {
		key   string
		value any
	}
	tests := []struct {
		name string
		c    *Context
		args args
	}{
		{
			name: "TestSet_0",
			c:    &Context{},
			args: args{
				key:   "key_0",
				value: "value_0",
			},
		},
		{
			name: "TestSet_1",
			c:    &Context{},
			args: args{
				key:   "key_1",
				value: 1,
			},
		},
		{
			name: "TestSet_2",
			c:    &Context{},
			args: args{
				key:   "key_2",
				value: true,
			},
		},
		{
			name: "TestSet_3",
			c:    &Context{},
			args: args{
				key:   "key_3",
				value: struct{}{},
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

			tt.c.Set(tt.args.key, tt.args.value)
			if tt.c.Keys[tt.args.key] != tt.args.value {
				t.Errorf("Expected %v, got %v", tt.args.value, tt.c.Keys[tt.args.key])
			}
		})
	}
}
