package httplib

import (
	"fmt"
	"testing"
)

func TestJSONBody_1(t *testing.T) {
	type testStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	testCases := []struct {
		name    string
		obj     testStruct
		wantErr bool
	}{
		{
			name: "valid object",
			obj: testStruct{
				Name: "John Doe",
				Age:  30,
			},
			wantErr: false,
		},
		{
			name: "invalid object",
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

			b := &BeegoHTTPRequest{}
			_, err := b.JSONBody(tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("JSONBody() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
