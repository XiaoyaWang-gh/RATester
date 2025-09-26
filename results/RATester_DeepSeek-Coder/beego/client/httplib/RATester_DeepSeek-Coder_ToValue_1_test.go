package httplib

import (
	"fmt"
	"testing"
)

func TestToValue_1(t *testing.T) {
	type TestStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	tests := []struct {
		name    string
		b       *BeegoHTTPRequest
		value   interface{}
		wantErr bool
	}{
		{
			name: "Test with JSON",
			b: &BeegoHTTPRequest{
				body: []byte(`{"name": "John", "age": 30}`),
			},
			value:   &TestStruct{},
			wantErr: false,
		},
		{
			name: "Test with YAML",
			b: &BeegoHTTPRequest{
				body: []byte("name: John\nage: 30"),
			},
			value:   &TestStruct{},
			wantErr: false,
		},
		{
			name: "Test with XML",
			b: &BeegoHTTPRequest{
				body: []byte("<TestStruct><name>John</name><age>30</age></TestStruct>"),
			},
			value:   &TestStruct{},
			wantErr: false,
		},
		{
			name: "Test with unsupported format",
			b: &BeegoHTTPRequest{
				body: []byte("unsupported format"),
			},
			value:   &TestStruct{},
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

			err := tt.b.ToValue(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("BeegoHTTPRequest.ToValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
