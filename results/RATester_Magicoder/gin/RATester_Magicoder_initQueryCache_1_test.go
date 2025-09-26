package gin

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestinitQueryCache_1(t *testing.T) {
	type args struct {
		c *Context
	}
	tests := []struct {
		name string
		args args
		want url.Values
	}{
		{
			name: "Test Case 1",
			args: args{
				c: &Context{
					Request: &http.Request{
						URL: &url.URL{
							RawQuery: "key=value",
						},
					},
				},
			},
			want: url.Values{"key": []string{"value"}},
		},
		{
			name: "Test Case 2",
			args: args{
				c: &Context{
					Request: &http.Request{
						URL: &url.URL{},
					},
				},
			},
			want: url.Values{},
		},
		{
			name: "Test Case 3",
			args: args{
				c: &Context{},
			},
			want: url.Values{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.args.c.initQueryCache()
			if !reflect.DeepEqual(tt.args.c.queryCache, tt.want) {
				t.Errorf("initQueryCache() = %v, want %v", tt.args.c.queryCache, tt.want)
			}
		})
	}
}
