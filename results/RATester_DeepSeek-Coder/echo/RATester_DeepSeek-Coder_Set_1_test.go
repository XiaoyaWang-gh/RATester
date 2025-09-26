package echo

import (
	"fmt"
	"testing"
)

func TestSet_1(t *testing.T) {
	type args struct {
		key string
		val interface{}
	}

	tests := []struct {
		name string
		c    *context
		args args
	}{
		{
			name: "TestSet_0",
			c:    &context{},
			args: args{
				key: "key",
				val: "value",
			},
		},
		{
			name: "TestSet_1",
			c:    &context{},
			args: args{
				key: "key",
				val: 123,
			},
		},
		{
			name: "TestSet_2",
			c:    &context{},
			args: args{
				key: "key",
				val: struct{}{},
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

			tt.c.Set(tt.args.key, tt.args.val)
			if got := tt.c.Get(tt.args.key); got != tt.args.val {
				t.Errorf("context.Set() = %v, want %v", got, tt.args.val)
			}
		})
	}
}
