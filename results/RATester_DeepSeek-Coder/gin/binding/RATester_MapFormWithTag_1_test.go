package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMapFormWithTag_1(t *testing.T) {
	type TestStruct struct {
		Name string `form:"name"`
		Age  int    `form:"age"`
	}

	testCases := []struct {
		name    string
		form    map[string][]string
		tag     string
		want    TestStruct
		wantErr bool
	}{
		{
			name: "valid form",
			form: map[string][]string{
				"name": {"John"},
				"age":  {"30"},
			},
			tag: "form",
			want: TestStruct{
				Name: "John",
				Age:  30,
			},
			wantErr: false,
		},
		{
			name: "invalid form",
			form: map[string][]string{
				"name": {"John"},
				"age":  {"not an int"},
			},
			tag:     "form",
			want:    TestStruct{},
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

			var got TestStruct
			err := MapFormWithTag(&got, tc.form, tc.tag)
			if (err != nil) != tc.wantErr {
				t.Errorf("MapFormWithTag() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("MapFormWithTag() got = %v, want %v", got, tc.want)
			}
		})
	}
}
