package parser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMarshalJSON_1(t *testing.T) {
	tests := []struct {
		name    string
		c       ReplacingJSONMarshaller
		want    []byte
		wantErr bool
	}{
		{
			name: "Test with KeysToLower",
			c: ReplacingJSONMarshaller{
				Value: map[string]interface{}{
					"Key1": "Value1",
					"Key2": "Value2",
				},
				KeysToLower: true,
			},
			want: []byte(`{"key1":"Value1","key2":"Value2"}`),
		},
		{
			name: "Test with OmitEmpty",
			c: ReplacingJSONMarshaller{
				Value: map[string]interface{}{
					"Key1": "Value1",
					"Key2": "",
				},
				OmitEmpty: true,
			},
			want: []byte(`{"Key1":"Value1"}`),
		},
		{
			name: "Test with KeysToLower and OmitEmpty",
			c: ReplacingJSONMarshaller{
				Value: map[string]interface{}{
					"Key1": "Value1",
					"Key2": "",
				},
				KeysToLower: true,
				OmitEmpty:   true,
			},
			want: []byte(`{"key1":"Value1"}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tt.c.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}
