package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetBoolField_1(t *testing.T) {

	type args struct {
		value string
		field reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "value is empty",
			args: args{
				value: "",
				field: reflect.ValueOf(false),
			},
			wantErr: false,
		},
		{
			name: "value is false",
			args: args{
				value: "false",
				field: reflect.ValueOf(false),
			},
			wantErr: false,
		},
		{
			name: "value is true",
			args: args{
				value: "true",
				field: reflect.ValueOf(false),
			},
			wantErr: false,
		},
		{
			name: "value is not bool",
			args: args{
				value: "not bool",
				field: reflect.ValueOf(false),
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

			if err := setBoolField(tt.args.value, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("setBoolField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
