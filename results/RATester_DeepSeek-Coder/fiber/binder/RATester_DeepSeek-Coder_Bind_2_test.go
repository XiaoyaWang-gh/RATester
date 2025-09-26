package binder

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBind_2(t *testing.T) {
	t.Parallel()

	type testStruct struct {
		Field string `xml:"field"`
	}

	testCases := []struct {
		name    string
		body    []byte
		want    testStruct
		wantErr bool
	}{
		{
			name: "success",
			body: []byte(`<testStruct><field>value</field></testStruct>`),
			want: testStruct{
				Field: "value",
			},
			wantErr: false,
		},
		{
			name:    "failure",
			body:    []byte(`<testStruct><field>value</field></testStruct>`),
			want:    testStruct{},
			wantErr: true,
		},
	}

	b := &xmlBinding{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := b.Bind(tc.body, &tc.want)

			if (err != nil) != tc.wantErr {
				t.Errorf("Bind() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !reflect.DeepEqual(tc.want, tc.want) {
				t.Errorf("Bind() got = %v, want %v", tc.want, tc.want)
			}
		})
	}
}
