package binding

import (
	"fmt"
	"testing"
)

func TestBindBody_5(t *testing.T) {
	b := jsonBinding{}

	type testStruct struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}

	testCases := []struct {
		name    string
		body    []byte
		obj     any
		wantErr bool
	}{
		{
			name: "valid json",
			body: []byte(`{"field1": "value1", "field2": 2}`),
			obj:  &testStruct{},
		},
		{
			name:    "invalid json",
			body:    []byte(`{"field1": "value1", "field2": "not an int"}`),
			obj:     &testStruct{},
			wantErr: true,
		},
		{
			name:    "nil obj",
			body:    []byte(`{"field1": "value1", "field2": 2}`),
			obj:     nil,
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

			err := b.BindBody(tc.body, tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("BindBody() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
