package hydratation

import (
	"fmt"
	"testing"
)

func TestHydrate_1(t *testing.T) {
	type TestStruct struct {
		Name string
		Age  int
	}

	testCases := []struct {
		name    string
		element interface{}
		wantErr bool
	}{
		{
			name: "Test case 1",
			element: &TestStruct{
				Name: "John Doe",
				Age:  30,
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			element: &TestStruct{
				Name: "",
				Age:  0,
			},
			wantErr: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := Hydrate(tt.element)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hydrate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
