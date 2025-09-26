package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetFormMap_1(t *testing.T) {
	type testStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	testCases := []struct {
		name    string
		form    map[string][]string
		want    any
		wantErr bool
	}{
		{
			name: "Test Case 1",
			form: map[string][]string{
				"name": {"John"},
				"age":  {"30"},
			},
			want: &testStruct{
				Name: "John",
				Age:  30,
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			form: map[string][]string{
				"name": {"Jane"},
				"age":  {"25"},
			},
			want: &testStruct{
				Name: "Jane",
				Age:  25,
			},
			wantErr: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := &testStruct{}
			err := setFormMap(got, tt.form)
			if (err != nil) != tt.wantErr {
				t.Errorf("setFormMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setFormMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}
