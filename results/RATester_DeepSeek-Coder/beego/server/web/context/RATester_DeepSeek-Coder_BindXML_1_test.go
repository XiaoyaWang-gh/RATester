package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBindXML_1(t *testing.T) {
	type testStruct struct {
		Name string `xml:"name"`
		Age  int    `xml:"age"`
	}

	testCases := []struct {
		name    string
		input   string
		want    testStruct
		wantErr bool
	}{
		{
			name:  "valid XML",
			input: `<testStruct><name>John</name><age>30</age></testStruct>`,
			want:  testStruct{Name: "John", Age: 30},
		},
		{
			name:    "invalid XML",
			input:   `<testStruct><name>John</name><age>30`,
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

			ctx := &Context{
				Input: &BeegoInput{
					RequestBody: []byte(tc.input),
				},
			}

			var got testStruct
			err := ctx.BindXML(&got)

			if (err != nil) != tc.wantErr {
				t.Errorf("BindXML() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !tc.wantErr && !reflect.DeepEqual(got, tc.want) {
				t.Errorf("BindXML() got = %v, want %v", got, tc.want)
			}
		})
	}
}
