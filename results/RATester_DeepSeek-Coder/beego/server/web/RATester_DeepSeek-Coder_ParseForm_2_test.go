package web

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestParseForm_2(t *testing.T) {
	type TestStruct struct {
		Name string `form:"name"`
		Age  int    `form:"age"`
	}

	testCases := []struct {
		name    string
		form    url.Values
		want    TestStruct
		wantErr bool
	}{
		{
			name: "valid form",
			form: url.Values{
				"name": []string{"John"},
				"age":  []string{"30"},
			},
			want: TestStruct{
				Name: "John",
				Age:  30,
			},
			wantErr: false,
		},
		{
			name: "invalid age",
			form: url.Values{
				"name": []string{"John"},
				"age":  []string{"not an int"},
			},
			wantErr: true,
		},
		{
			name: "missing age",
			form: url.Values{
				"name": []string{"John"},
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
			err := ParseForm(tc.form, &got)
			if (err != nil) != tc.wantErr {
				t.Errorf("ParseForm() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ParseForm() got = %v, want %v", got, tc.want)
			}
		})
	}
}
