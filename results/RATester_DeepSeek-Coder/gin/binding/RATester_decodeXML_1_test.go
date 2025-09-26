package binding

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestDecodeXML_1(t *testing.T) {
	type TestStruct struct {
		Field string `xml:"field"`
	}

	testCases := []struct {
		name    string
		input   string
		want    TestStruct
		wantErr bool
	}{
		{
			name:  "valid xml",
			input: `<TestStruct><field>value</field></TestStruct>`,
			want:  TestStruct{Field: "value"},
		},
		{
			name:    "invalid xml",
			input:   `<TestStruct><field>value`,
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

			r := strings.NewReader(tc.input)
			obj := &TestStruct{}
			err := decodeXML(r, obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("decodeXML() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(*obj, tc.want) {
				t.Errorf("decodeXML() = %v, want %v", *obj, tc.want)
			}
		})
	}
}
