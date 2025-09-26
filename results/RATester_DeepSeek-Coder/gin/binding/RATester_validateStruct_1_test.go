package binding

import (
	"fmt"
	"testing"
)

func TestValidateStruct_1(t *testing.T) {
	type testStruct struct {
		Field1 string `validate:"required"`
		Field2 int    `validate:"gt=0"`
	}

	testCases := []struct {
		name    string
		obj     any
		wantErr bool
	}{
		{
			name: "valid",
			obj: &testStruct{
				Field1: "test",
				Field2: 1,
			},
			wantErr: false,
		},
		{
			name: "invalid",
			obj: &testStruct{
				Field1: "",
				Field2: 0,
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

			v := &defaultValidator{}
			err := v.validateStruct(tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("validateStruct() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
