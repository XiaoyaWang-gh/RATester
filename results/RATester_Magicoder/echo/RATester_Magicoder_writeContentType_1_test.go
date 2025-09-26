package echo

import (
	"fmt"
	"testing"
)

func TestWriteContentType_1(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		c    *context
		args args
		want string
	}{
		{
			name: "Test Case 1",
			c:    &context{},
			args: args{value: "text/plain"},
			want: "text/plain",
		},
		{
			name: "Test Case 2",
			c:    &context{},
			args: args{value: "text/html"},
			want: "text/html",
		},
		{
			name: "Test Case 3",
			c:    &context{},
			args: args{value: "application/json"},
			want: "application/json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.c.writeContentType(tt.args.value)
			header := tt.c.Response().Header()
			if got := header.Get(HeaderContentType); got != tt.want {
				t.Errorf("writeContentType() = %v, want %v", got, tt.want)
			}
		})
	}
}
