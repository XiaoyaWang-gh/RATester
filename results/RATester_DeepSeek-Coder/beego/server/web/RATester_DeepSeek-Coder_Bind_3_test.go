package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestBind_3(t *testing.T) {
	type testStruct struct {
		Name string `form:"name"`
		Age  int    `form:"age"`
	}

	testCases := []struct {
		name    string
		obj     testStruct
		wantErr bool
	}{
		{
			name: "valid data",
			obj: testStruct{
				Name: "John",
				Age:  30,
			},
			wantErr: false,
		},
		{
			name: "invalid data",
			obj: testStruct{
				Name: "",
				Age:  0,
			},
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

			c := &Controller{
				Ctx: &context.Context{},
			}

			err := c.Bind(&tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("Bind() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
