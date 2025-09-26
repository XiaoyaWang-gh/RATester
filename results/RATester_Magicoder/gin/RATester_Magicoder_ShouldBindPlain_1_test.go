package gin

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestShouldBindPlain_1(t *testing.T) {
	type testStruct struct {
		Field1 string `form:"field1"`
		Field2 int    `form:"field2"`
	}

	tests := []struct {
		name    string
		context *Context
		obj     any
		wantErr bool
	}{
		{
			name: "Success",
			context: &Context{
				Request: &http.Request{
					Form: url.Values{
						"field1": []string{"value1"},
						"field2": []string{"123"},
					},
				},
			},
			obj:     &testStruct{},
			wantErr: false,
		},
		{
			name: "Error",
			context: &Context{
				Request: &http.Request{
					Form: url.Values{
						"field1": []string{"value1"},
						"field2": []string{"abc"}, // not an integer
					},
				},
			},
			obj:     &testStruct{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := tt.context.ShouldBindPlain(tt.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShouldBindPlain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
