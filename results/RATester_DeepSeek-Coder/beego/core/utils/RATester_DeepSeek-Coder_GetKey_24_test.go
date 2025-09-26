package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetKey_24(t *testing.T) {
	type fields struct {
		Key   interface{}
		Value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "TestGetKey_1",
			fields: fields{
				Key:   "testKey",
				Value: "testValue",
			},
			want: "testKey",
		},
		{
			name: "TestGetKey_2",
			fields: fields{
				Key:   123,
				Value: "testValue",
			},
			want: 123,
		},
		{
			name: "TestGetKey_3",
			fields: fields{
				Key:   true,
				Value: "testValue",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &SimpleKV{
				Key:   tt.fields.Key,
				Value: tt.fields.Value,
			}
			if got := s.GetKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
