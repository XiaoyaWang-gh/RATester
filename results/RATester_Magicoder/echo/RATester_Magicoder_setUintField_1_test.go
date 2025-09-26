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
			name: "Test with valid input",
			args: args{
				value:   "10",
				bitSize: 64,
				field:   reflect.ValueOf(uint64(0)),
			},
			wantErr: false,
		},
		{
			name: "Test with invalid input",
			args: args{
				value:   "invalid",
				bitSize: 64,
				field:   reflect.ValueOf(uint64(0)),
			},
			wantErr: true,
		},
		{
			name: "Test with empty input",
			args: args{
				value:   "",
				bitSize: 64,
				field:   reflect.ValueOf(uint64(0)),
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
