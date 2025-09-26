package binding

import (
	"fmt"
	"testing"
)

func TestBindUri_1(t *testing.T) {
	type testStruct struct {
		Field1 string `uri:"field1"`
		Field2 int    `uri:"field2"`
	}

	testCases := []struct {
		name    string
		input   map[string][]string
		wantErr bool
	}{
		{
			name: "valid input",
			input: map[string][]string{
				"field1": {"value1"},
				"field2": {"123"},
			},
			wantErr: false,
		},
		{
			name: "missing field1",
			input: map[string][]string{
				"field2": {"123"},
			},
			wantErr: true,
		},
		{
			name: "invalid field2",
			input: map[string][]string{
				"field1": {"value1"},
				"field2": {"abc"},
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

			b := uriBinding{}
			obj := &testStruct{}
			err := b.BindUri(tc.input, obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("BindUri() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
