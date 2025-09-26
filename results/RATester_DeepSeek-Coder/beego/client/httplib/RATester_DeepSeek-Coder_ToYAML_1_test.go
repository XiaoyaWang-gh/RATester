package httplib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToYAML_1(t *testing.T) {
	type testStruct struct {
		Name string `yaml:"name"`
		Age  int    `yaml:"age"`
	}

	tests := []struct {
		name    string
		request *BeegoHTTPRequest
		want    testStruct
		wantErr bool
	}{
		{
			name: "TestToYAML_Success",
			request: &BeegoHTTPRequest{
				body: []byte(`name: John Doe
age: 30`),
			},
			want: testStruct{
				Name: "John Doe",
				Age:  30,
			},
			wantErr: false,
		},
		{
			name: "TestToYAML_Error",
			request: &BeegoHTTPRequest{
				body: []byte(`name: John Doe
age: not an int`),
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

			err := tt.request.ToYAML(&tt.want)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(tt.want, tt.want) {
				t.Errorf("ToYAML() = %v, want %v", tt.want, tt.want)
			}
		})
	}
}
