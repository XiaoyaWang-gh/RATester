package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetFloatField_1(t *testing.T) {
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
			name: "Test with valid float value",
			args: args{
				value:   "1.23",
				bitSize: 64,
				field:   reflect.ValueOf(float64(0)),
			},
			wantErr: false,
		},
		{
			name: "Test with empty string",
			args: args{
				value:   "",
				bitSize: 64,
				field:   reflect.ValueOf(float64(0)),
			},
			wantErr: false,
		},
		{
			name: "Test with invalid float value",
			args: args{
				value:   "invalid",
				bitSize: 64,
				field:   reflect.ValueOf(float64(0)),
			},
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

			if err := setFloatField(tt.args.value, tt.args.bitSize, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("setFloatField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
