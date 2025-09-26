package echo

import (
	"fmt"
	"testing"
)

func TestInsert_1(t *testing.T) {
	type args struct {
		method string
		path   string
		h      HandlerFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Case 1",
			args: args{
				method: "GET",
				path:   "/test",
				h:      func(c Context) error { return nil },
			},
		},
		{
			name: "Test Case 2",
			args: args{
				method: "POST",
				path:   "/test/:id",
				h:      func(c Context) error { return nil },
			},
		},
		{
			name: "Test Case 3",
			args: args{
				method: "PUT",
				path:   "/test/*",
				h:      func(c Context) error { return nil },
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

			r := &Router{}
			r.insert(tt.args.method, tt.args.path, tt.args.h)
		})
	}
}
