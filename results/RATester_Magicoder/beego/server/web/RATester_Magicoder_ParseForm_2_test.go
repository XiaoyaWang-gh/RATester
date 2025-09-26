package web

import (
	"fmt"
	"net/url"
	"testing"
)

func TestParseForm_2(t *testing.T) {
	type TestStruct struct {
		Field1 string `form:"field1"`
		Field2 int    `form:"field2"`
	}

	testCases := []struct {
		name    string
		form    url.Values
		obj     interface{}
		wantErr bool
	}{
		{
			name: "valid form",
			form: url.Values{
				"field1": []string{"value1"},
				"field2": []string{"123"},
			},
			obj: &TestStruct{},
		},
		{
			name: "invalid form",
			form: url.Values{
				"field1": []string{"value1"},
				"field2": []string{"abc"}, // not an integer
			},
			obj:     &TestStruct{},
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

			err := ParseForm(tc.form, tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("ParseForm() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
