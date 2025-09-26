package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMapForm_1(t *testing.T) {
	type TestStruct struct {
		Name string `form:"name"`
		Age  int    `form:"age"`
	}

	testCases := []struct {
		name    string
		form    map[string][]string
		want    TestStruct
		wantErr bool
	}{
		{
			name: "valid form",
			form: map[string][]string{
				"name": {"John"},
				"age":  {"30"},
			},
			want: TestStruct{
				Name: "John",
				Age:  30,
			},
			wantErr: false,
		},
		{
			name: "missing age",
			form: map[string][]string{
				"name": {"John"},
			},
			want: TestStruct{
				Name: "John",
				Age:  0,
			},
			wantErr: false,
		},
		{
			name: "invalid age",
			form: map[string][]string{
				"name": {"John"},
				"age":  {"not an int"},
			},
			want: TestStruct{
				Name: "John",
				Age:  0,
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

			var got TestStruct
			err := mapForm(&got, tc.form)
			if (err != nil) != tc.wantErr {
				t.Errorf("mapForm() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("mapForm() got = %v, want %v", got, tc.want)
			}
		})
	}
}
