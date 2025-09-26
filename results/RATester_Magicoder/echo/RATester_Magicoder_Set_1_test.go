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
				key: "key_0",
				val: "val_0",
			},
		},
		{
			name: "TestSet_1",
			c:    &context{},
			args: args{
				key: "key_1",
				val: 1,
			},
		},
		{
			name: "TestSet_2",
			c:    &context{},
			args: args{
				key: "key_2",
				val: true,
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
		})
	}
}
