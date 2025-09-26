package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBindXML_1(t *testing.T) {
	type TestStruct struct {
		Field1 string `xml:"field1"`
		Field2 int    `xml:"field2"`
	}

	testCases := []struct {
		name    string
		input   string
		want    TestStruct
		wantErr bool
	}{
		{
			name:  "Valid XML",
			input: `<TestStruct><field1>value1</field1><field2>123</field2></TestStruct>`,
			want: TestStruct{
				Field1: "value1",
				Field2: 123,
			},
			wantErr: false,
		},
		{
			name:    "Invalid XML",
			input:   `<TestStruct><field1>value1</field1><field2>abc</field2></TestStruct>`,
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

			ctx := &Context{
				Input: &BeegoInput{
					RequestBody: []byte(tc.input),
				},
			}

			var got TestStruct
			err := ctx.BindXML(&got)

			if (err != nil) != tc.wantErr {
				t.Errorf("BindXML() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("BindXML() = %v, want %v", got, tc.want)
			}
		})
	}
}
