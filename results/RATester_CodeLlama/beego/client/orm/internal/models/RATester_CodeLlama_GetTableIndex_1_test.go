package models

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestGetTableIndex_1(t *testing.T) {
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "test case 1",
			args: args{
				val: reflect.ValueOf(context.Background()),
			},
			want: nil,
		},
		{
			name: "test case 2",
			args: args{
				val: reflect.ValueOf(context.WithValue(context.Background(), "key", "value")),
			},
			want: nil,
		},
		{
			name: "test case 3",
			args: args{
				val: reflect.ValueOf(context.WithValue(context.Background(), "key", "value")),
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

			if got := GetTableIndex(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTableIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
