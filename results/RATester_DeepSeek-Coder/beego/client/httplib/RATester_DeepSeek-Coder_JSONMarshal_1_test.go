package httplib

import (
	"fmt"
	"testing"
)

func TestJSONMarshal_1(t *testing.T) {
	type testStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	testCases := []struct {
		name    string
		input   testStruct
		want    string
		wantErr bool
	}{
		{
			name: "Test case 1",
			input: testStruct{
				Name: "John",
				Age:  30,
			},
			want:    `{"name":"John","age":30}`,
			wantErr: false,
		},
		{
			name: "Test case 2",
			input: testStruct{
				Name: "Jane",
				Age:  25,
			},
			want:    `{"name":"Jane","age":25}`,
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			b := &BeegoHTTPRequest{}
			got, err := b.JSONMarshal(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("JSONMarshal() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if string(got) != tc.want {
				t.Errorf("JSONMarshal() = %v, want %v", string(got), tc.want)
			}
		})
	}
}
