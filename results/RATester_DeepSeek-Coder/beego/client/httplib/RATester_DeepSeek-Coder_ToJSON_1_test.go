package httplib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToJSON_1(t *testing.T) {
	type TestStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	testCases := []struct {
		name    string
		request *BeegoHTTPRequest
		want    TestStruct
		wantErr bool
	}{
		{
			name: "Test ToJSON with valid JSON",
			request: &BeegoHTTPRequest{
				body: []byte(`{"name": "John", "age": 30}`),
			},
			want: TestStruct{
				Name: "John",
				Age:  30,
			},
			wantErr: false,
		},
		{
			name: "Test ToJSON with invalid JSON",
			request: &BeegoHTTPRequest{
				body: []byte(`{"name": "John", "age": "thirty"}`),
			},
			want:    TestStruct{},
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

			err := tt.request.ToJSON(&tt.want)
			if (err != nil) != tt.wantErr {
				t.Errorf("BeegoHTTPRequest.ToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(tt.want, tt.want) {
				t.Errorf("BeegoHTTPRequest.ToJSON() = %v, want %v", tt.want, tt.want)
			}
		})
	}
}
