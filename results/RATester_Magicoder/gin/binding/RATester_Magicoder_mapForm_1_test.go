package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestmapForm_1(t *testing.T) {
	type TestStruct struct {
		Field1 string `form:"field1"`
		Field2 int    `form:"field2"`
	}

	testCases := []struct {
		name    string
		form    map[string][]string
		want    TestStruct
		wantErr bool
	}{
		{
			name: "Happy path",
			form: map[string][]string{
				"field1": {"value1"},
				"field2": {"123"},
			},
			want: TestStruct{
				Field1: "value1",
				Field2: 123,
			},
			wantErr: false,
		},
		{
			name: "Missing field",
			form: map[string][]string{
				"field1": {"value1"},
			},
			want: TestStruct{
				Field1: "value1",
			},
			wantErr: false,
		},
		{
			name: "Invalid field",
			form: map[string][]string{
				"field1": {"value1"},
				"field2": {"not_an_int"},
			},
			want: TestStruct{
				Field1: "value1",
			},
			wantErr: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := TestStruct{}
			err := mapForm(&got, tt.form)
			if (err != nil) != tt.wantErr {
				t.Errorf("mapForm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapForm() = %v, want %v", got, tt.want)
			}
		})
	}
}
