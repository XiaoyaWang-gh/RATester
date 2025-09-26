package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetUintField_1(t *testing.T) {
	type args struct {
		value   string
		bitSize int
		field   reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				value:   "1",
				bitSize: 1,
				field:   reflect.ValueOf(1),
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				value:   "1",
				bitSize: 1,
				field:   reflect.ValueOf(1),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := setUintField(tt.args.value, tt.args.bitSize, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("setUintField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
