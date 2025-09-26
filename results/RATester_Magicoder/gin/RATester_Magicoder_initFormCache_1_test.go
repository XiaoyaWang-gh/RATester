package gin

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestinitFormCache_1(t *testing.T) {
	type args struct {
		c *Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Case 1",
			args: args{
				c: &Context{
					Request: &http.Request{
						PostForm: url.Values{
							"key": []string{"value"},
						},
					},
				},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				c: &Context{
					Request: &http.Request{
						PostForm: url.Values{},
					},
				},
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

			tt.args.c.initFormCache()
		})
	}
}
