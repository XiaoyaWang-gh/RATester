package gin

import (
	"fmt"
	"testing"
)

func TestBindTOML_1(t *testing.T) {
	type testStruct struct {
		Name string `toml:"name"`
		Age  int    `toml:"age"`
	}

	testCases := []struct {
		name    string
		context *Context
		obj     any
		wantErr bool
	}{
		{
			name:    "success",
			context: &Context{},
			obj:     &testStruct{},
			wantErr: false,
		},
		{
			name:    "failure",
			context: &Context{},
			obj:     "invalid",
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

			err := tc.context.BindTOML(tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("BindTOML() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
