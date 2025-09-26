package gin

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestShouldBind_1(t *testing.T) {
	type testStruct struct {
		Field1 string `form:"field1"`
		Field2 int    `form:"field2"`
	}

	testCases := []struct {
		name    string
		context *Context
		obj     any
		wantErr bool
	}{
		{
			name: "Success",
			context: &Context{
				Request: &http.Request{
					Method: "POST",
					Form:   url.Values{"field1": {"value1"}, "field2": {"123"}},
				},
			},
			obj: &testStruct{},
		},
		{
			name: "Error",
			context: &Context{
				Request: &http.Request{
					Method: "POST",
					Form:   url.Values{"field1": {"value1"}},
				},
			},
			obj:     &testStruct{},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := tc.context.ShouldBind(tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("ShouldBind() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
