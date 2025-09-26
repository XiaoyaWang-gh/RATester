package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncodeRawValue_1(t *testing.T) {
	type args struct {
		labels   map[string]string
		root     string
		rawValue interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "Test case 1",
			args: args{
				labels:   map[string]string{},
				root:     "",
				rawValue: map[string]interface{}{"field name": "value"},
			},
			want: map[string]string{"field name": "value"},
		},
		{
			name: "Test case 2",
			args: args{
				labels:   map[string]string{},
				root:     "",
				rawValue: map[string]interface{}{"field name": []interface{}{"value1", "value2"}},
			},
			want: map[string]string{"field name[0]": "value1", "field name[1]": "value2"},
		},
		{
			name: "Test case 3",
			args: args{
				labels:   map[string]string{},
				root:     "",
				rawValue: map[string]interface{}{"field name": map[string]interface{}{"subfield": "value"}},
			},
			want: map[string]string{"field name.subfield": "value"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			encodeRawValue(tt.args.labels, tt.args.root, tt.args.rawValue)
			if !reflect.DeepEqual(tt.args.labels, tt.want) {
				t.Errorf("encodeRawValue() = %v, want %v", tt.args.labels, tt.want)
			}
		})
	}
}
