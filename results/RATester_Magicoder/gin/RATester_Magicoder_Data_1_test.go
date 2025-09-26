package gin

import (
	"fmt"
	"testing"
)

func TestData_1(t *testing.T) {
	type args struct {
		code        int
		contentType string
		data        []byte
	}
	tests := []struct {
		name string
		c    *Context
		args args
	}{
		{
			name: "Test Case 1",
			c:    &Context{},
			args: args{
				code:        200,
				contentType: "application/json",
				data:        []byte(`{"key": "value"}`),
			},
		},
		{
			name: "Test Case 2",
			c:    &Context{},
			args: args{
				code:        404,
				contentType: "text/plain",
				data:        []byte("Not Found"),
			},
		},
		{
			name: "Test Case 3",
			c:    &Context{},
			args: args{
				code:        500,
				contentType: "application/xml",
				data:        []byte(`<error>Internal Server Error</error>`),
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

			tt.c.Data(tt.args.code, tt.args.contentType, tt.args.data)
		})
	}
}
