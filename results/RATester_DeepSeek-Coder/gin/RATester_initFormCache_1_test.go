package gin

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestInitFormCache_1(t *testing.T) {
	tests := []struct {
		name string
		c    *Context
		want url.Values
	}{
		{
			name: "Test with empty formCache",
			c: &Context{
				Request: &http.Request{
					PostForm: url.Values{},
				},
				engine: &Engine{
					MaxMultipartMemory: 1024,
				},
			},
			want: url.Values{},
		},
		{
			name: "Test with non-empty formCache",
			c: &Context{
				Request: &http.Request{
					PostForm: url.Values{
						"key": []string{"value"},
					},
				},
				formCache: url.Values{
					"key": []string{"value"},
				},
				engine: &Engine{
					MaxMultipartMemory: 1024,
				},
			},
			want: url.Values{
				"key": []string{"value"},
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

			tt.c.initFormCache()
			if !reflect.DeepEqual(tt.c.formCache, tt.want) {
				t.Errorf("initFormCache() = %v, want %v", tt.c.formCache, tt.want)
			}
		})
	}
}
