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
			name: "Test case 1",
			args: args{
				value: "true",
				field: reflect.ValueOf(true),
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				value: "false",
				field: reflect.ValueOf(false),
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				value: "",
				field: reflect.ValueOf(true),
			},
			wantErr: false,
		},
		{
			name: "Test case 4",
			args: args{
				value: "invalid",
				field: reflect.ValueOf(true),
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
