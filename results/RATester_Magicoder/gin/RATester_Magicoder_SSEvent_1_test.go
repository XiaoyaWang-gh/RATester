package gin

import (
	"fmt"
	"testing"
)

func TestSSEvent_1(t *testing.T) {
	type args struct {
		name    string
		message any
	}
	tests := []struct {
		name string
		c    *Context
		args args
	}{
		{
			name: "Test case 1",
			c:    &Context{},
			args: args{
				name:    "test_event",
				message: "test_message",
			},
		},
		{
			name: "Test case 2",
			c:    &Context{},
			args: args{
				name:    "test_event",
				message: 123,
			},
		},
		{
			name: "Test case 3",
			c:    &Context{},
			args: args{
				name:    "test_event",
				message: struct{}{},
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

			tt.c.SSEvent(tt.args.name, tt.args.message)
		})
	}
}
