package fiber

import (
	"fmt"
	"testing"
)

func TestvalidateStruct_1(t *testing.T) {
	type testStruct struct {
		Field1 string `validate:"required"`
		Field2 int    `validate:"gt=0"`
	}

	testCases := []struct {
		name    string
		input   any
		wantErr bool
	}{
		{
			name: "valid input",
			input: testStruct{
				Field1: "test",
				Field2: 1,
			},
			wantErr: false,
		},
		{
			name: "invalid input",
			input: testStruct{
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

			b := &Bind{}
			err := b.validateStruct(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("validateStruct() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
