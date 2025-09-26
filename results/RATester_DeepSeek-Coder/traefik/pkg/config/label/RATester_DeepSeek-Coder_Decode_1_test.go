package label

import (
	"fmt"
	"testing"
)

func TestDecode_1(t *testing.T) {
	type TestStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	testCases := []struct {
		name    string
		labels  map[string]string
		element interface{}
		filters []string
		wantErr bool
	}{
		{
			name: "valid input",
			labels: map[string]string{
				"name": "John",
				"age":  "30",
			},
			element: &TestStruct{},
			filters: []string{"name", "age"},
			wantErr: false,
		},
		{
			name: "invalid input",
			labels: map[string]string{
				"name": "John",
				"age":  "not_an_int",
			},
			element: &TestStruct{},
			filters: []string{"name", "age"},
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

			err := Decode(tc.labels, tc.element, tc.filters...)
			if (err != nil) != tc.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
