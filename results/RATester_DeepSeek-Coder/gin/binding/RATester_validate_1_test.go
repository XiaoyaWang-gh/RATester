package binding

import (
	"fmt"
	"testing"
)

func TestValidate_1(t *testing.T) {
	type testStruct struct {
		Name string `validate:"required"`
		Age  int    `validate:"gte=0,lte=130"`
	}

	testCases := []struct {
		name    string
		input   any
		wantErr bool
	}{
		{
			name: "valid",
			input: testStruct{
				Name: "John Doe",
				Age:  25,
			},
			wantErr: false,
		},
		{
			name: "missing name",
			input: testStruct{
				Age: 25,
			},
			wantErr: true,
		},
		{
			name: "invalid age",
			input: testStruct{
				Name: "John Doe",
				Age:  131,
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

			err := validate(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
