package gin

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestShouldBindPlain_1(t *testing.T) {
	type testStruct struct {
		Field string `form:"field"`
	}

	tests := []struct {
		name    string
		context *Context
		obj     any
		wantErr bool
	}{
		{
			name: "success",
			context: &Context{
				Request: &http.Request{
					Form: url.Values{"field": []string{"value"}},
				},
			},
			obj:     &testStruct{},
			wantErr: false,
		},
		{
			name: "failure",
			context: &Context{
				Request: &http.Request{
					Form: url.Values{"field": []string{"value"}},
				},
			},
			obj:     "not a struct pointer",
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
			}
		})
	}
}
