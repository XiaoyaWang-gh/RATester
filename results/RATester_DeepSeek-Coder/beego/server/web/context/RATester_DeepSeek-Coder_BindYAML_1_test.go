package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBindYAML_1(t *testing.T) {
	type testStruct struct {
		Name string `yaml:"name"`
		Age  int    `yaml:"age"`
	}

	testCases := []struct {
		name    string
		input   string
		want    testStruct
		wantErr bool
	}{
		{
			name:  "valid yaml",
			input: "name: John\nage: 30\n",
			want:  testStruct{Name: "John", Age: 30},
		},
		{
			name:    "invalid yaml",
			input:   "name: John\nage: not_an_int\n",
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
			err := ctx.BindYAML(&got)

			if (err != nil) != tc.wantErr {
				t.Errorf("BindYAML() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !tc.wantErr && !reflect.DeepEqual(got, tc.want) {
				t.Errorf("BindYAML() got = %v, want %v", got, tc.want)
			}
		})
	}
}
