package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBindBody_7(t *testing.T) {
	type testStruct struct {
		Name string `yaml:"name"`
		Age  int    `yaml:"age"`
	}

	testCases := []struct {
		name    string
		body    string
		want    testStruct
		wantErr bool
	}{
		{
			name: "valid yaml",
			body: "name: John\nage: 30\n",
			want: testStruct{
				Name: "John",
				Age:  30,
			},
			wantErr: false,
		},
		{
			name:    "invalid yaml",
			body:    "name: John\nage: not_an_int\n",
			want:    testStruct{},
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

			yb := yamlBinding{}
			got := testStruct{}
			err := yb.BindBody([]byte(tc.body), &got)
			if (err != nil) != tc.wantErr {
				t.Errorf("BindBody() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("BindBody() got = %v, want %v", got, tc.want)
			}
		})
	}
}
