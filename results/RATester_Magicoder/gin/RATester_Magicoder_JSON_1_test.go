package gin

import (
	"fmt"
	"testing"
)

func TestJSON_1(t *testing.T) {
	type args struct {
		code int
		obj  any
	}
	tests := []struct {
		name string
		c    *Context
		args args
	}{
		{
			name: "TestJSON_0",
			c:    &Context{},
			args: args{
				code: 200,
				obj:  "test",
			},
		},
		{
			name: "TestJSON_1",
			c:    &Context{},
			args: args{
				code: 404,
				obj:  "not found",
			},
		},
		{
			name: "TestJSON_2",
			c:    &Context{},
			args: args{
				code: 500,
				obj:  "internal server error",
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

			tt.c.JSON(tt.args.code, tt.args.obj)
		})
	}
}
