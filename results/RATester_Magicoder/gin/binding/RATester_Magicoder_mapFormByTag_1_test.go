package binding

import (
	"fmt"
	"testing"
)

func TestmapFormByTag_1(t *testing.T) {
	type TestStruct struct {
		Field1 string `form:"field1"`
		Field2 int    `form:"field2"`
	}

	testCases := []struct {
		name    string
		ptr     any
		form    map[string][]string
		tag     string
		wantErr bool
	}{
		{
			name: "Test with valid input",
			ptr:  &TestStruct{},
			form: map[string][]string{
				"field1": {"value1"},
				"field2": {"10"},
			},
			tag:     "form",
			wantErr: false,
		},
		{
			name: "Test with invalid input",
			ptr:  &TestStruct{},
			form: map[string][]string{
				"field1": {"value1"},
				"field2": {"invalid"},
			},
			tag:     "form",
			wantErr: true,
		},
		{
			name: "Test with nil input",
			ptr:  nil,
			form: map[string][]string{
				"field1": {"value1"},
				"field2": {"10"},
			},
			tag:     "form",
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

			err := mapFormByTag(tc.ptr, tc.form, tc.tag)
			if (err != nil) != tc.wantErr {
				t.Errorf("mapFormByTag() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
