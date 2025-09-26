package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetValue_3(t *testing.T) {
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
			name: "TestGetValue_0",
			fields: fields{
				Key:   "Key_0",
				Value: "Value_0",
			},
			want: "Value_0",
		},
		{
			name: "TestGetValue_1",
			fields: fields{
				Key:   1,
				Value: 100,
			},
			want: 100,
		},
		{
			name: "TestGetValue_2",
			fields: fields{
				Key:   "Key_2",
				Value: nil,
			},
			want: nil,
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
			if got := s.GetValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
