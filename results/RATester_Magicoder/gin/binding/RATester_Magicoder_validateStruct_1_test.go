package binding

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestvalidateStruct_1(t *testing.T) {
	v := &defaultValidator{
		validate: validator.New(),
	}

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
				Age:  30,
			},
			wantErr: false,
		},
		{
			name: "invalid",
			input: testStruct{
				Name: "",
				Age:  -1,
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

			err := v.validateStruct(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("validateStruct() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
