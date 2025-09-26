package context

import (
	"fmt"
	"testing"
)

func TestBindJSON_1(t *testing.T) {
	type TestStruct struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}

	testCases := []struct {
		name    string
		context *Context
		obj     interface{}
		wantErr bool
	}{
		{
			name: "Success",
			context: &Context{
				Input: &BeegoInput{
					RequestBody: []byte(`{"field1": "value1", "field2": 123}`),
				},
			},
			obj: &TestStruct{},
		},
		{
			name: "Invalid JSON",
			context: &Context{
				Input: &BeegoInput{
					RequestBody: []byte(`{field1": "value1", "field2": 123}`),
				},
			},
			obj:     &TestStruct{},
			wantErr: true,
		},
		{
			name: "Invalid Object",
			context: &Context{
				Input: &BeegoInput{
					RequestBody: []byte(`{"field1": "value1", "field2": 123}`),
				},
			},
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

			err := tc.context.BindJSON(tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("BindJSON() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
