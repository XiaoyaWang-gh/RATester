package label

import (
	"fmt"
	"testing"
)

func TestDecode_1(t *testing.T) {
	type TestStruct struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}

	testCases := []struct {
		name    string
		labels  map[string]string
		element interface{}
		filters []string
		wantErr bool
	}{
		{
			name: "Test case 1",
			labels: map[string]string{
				"field1": "value1",
				"field2": "10",
			},
			element: &TestStruct{},
			filters: []string{"field1"},
			wantErr: false,
		},
		{
			name: "Test case 2",
			labels: map[string]string{
				"field1": "value1",
				"field2": "not_an_int",
			},
			element: &TestStruct{},
			filters: []string{"field1", "field2"},
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
				return
			}
		})
	}
}
