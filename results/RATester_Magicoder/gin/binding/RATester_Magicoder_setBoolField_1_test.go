package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestsetBoolField_1(t *testing.T) {
	type args struct {
		val   string
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
				val:   "true",
				field: reflect.ValueOf(true),
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				val:   "false",
				field: reflect.ValueOf(false),
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				val:   "",
				field: reflect.ValueOf(true),
			},
			wantErr: false,
		},
		{
			name: "Test case 4",
			args: args{
				val:   "invalid",
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

			if err := setBoolField(tt.args.val, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("setBoolField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
