package gin

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestShouldBind_1(t *testing.T) {
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
			name: "should bind successfully",
			context: &Context{
				Request: &http.Request{
					Method: "POST",
					Form:   url.Values{"field": []string{"test"}},
				},
			},
			obj:     &testStruct{},
			wantErr: false,
		},
		{
			name: "should bind with empty form",
			context: &Context{
				Request: &http.Request{
					Method: "POST",
					Form:   url.Values{},
				},
			},
			obj:     &testStruct{},
			wantErr: false,
		},
		{
			name: "should bind with nil form",
			context: &Context{
				Request: &http.Request{
					Method: "POST",
					Form:   nil,
				},
			},
			obj:     &testStruct{},
			wantErr: false,
		},
		{
			name: "should bind with invalid method",
			context: &Context{
				Request: &http.Request{
					Method: "GET",
					Form:   url.Values{"field": []string{"test"}},
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

			err := tt.context.ShouldBind(tt.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShouldBind() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
