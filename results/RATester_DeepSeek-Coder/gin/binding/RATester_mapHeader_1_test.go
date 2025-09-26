package binding

import (
	"fmt"
	"testing"
)

func TestMapHeader_1(t *testing.T) {
	type TestStruct struct {
		Name string `header:"name"`
		Age  int    `header:"age"`
	}

	testCases := []struct {
		name    string
		headers map[string][]string
		wantErr bool
	}{
		{
			name: "valid headers",
			headers: map[string][]string{
				"name": {"John Doe"},
				"age":  {"30"},
			},
			wantErr: false,
		},
		{
			name: "missing headers",
			headers: map[string][]string{
				"name": {},
				"age":  {},
			},
			wantErr: true,
		},
		{
			name: "invalid headers",
			headers: map[string][]string{
				"name": {"John Doe"},
				"age":  {"not an int"},
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

			ts := &TestStruct{}
			err := mapHeader(ts, tc.headers)
			if (err != nil) != tc.wantErr {
				t.Errorf("mapHeader() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
