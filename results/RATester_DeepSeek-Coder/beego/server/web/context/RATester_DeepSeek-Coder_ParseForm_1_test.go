package context

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestParseForm_1(t *testing.T) {
	type testStruct struct {
		Name string `form:"name"`
		Age  int    `form:"age"`
	}

	tests := []struct {
		name    string
		form    url.Values
		want    testStruct
		wantErr bool
	}{
		{
			name: "valid form",
			form: url.Values{
				"name": []string{"John"},
				"age":  []string{"30"},
			},
			want: testStruct{
				Name: "John",
				Age:  30,
			},
			wantErr: false,
		},
		{
			name: "invalid form",
			form: url.Values{
				"name": []string{"John"},
				"age":  []string{"not an int"},
			},
			want:    testStruct{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			var got testStruct
			err := ParseForm(tt.form, &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseForm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseForm() got = %v, want %v", got, tt.want)
			}
		})
	}
}
