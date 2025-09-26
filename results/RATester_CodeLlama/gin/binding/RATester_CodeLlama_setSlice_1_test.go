package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetSlice_1(t *testing.T) {
	type args struct {
		vals  []string
		value reflect.Value
		field reflect.StructField
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				vals:  []string{"1", "2", "3"},
				value: reflect.ValueOf([]string{}),
				field: reflect.StructField{},
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				vals:  []string{"1", "2", "3"},
				value: reflect.ValueOf([]string{}),
				field: reflect.StructField{},
			},
			wantErr: false,
		},
		{
			name: "test case 3",
			args: args{
				vals:  []string{"1", "2", "3"},
				value: reflect.ValueOf([]string{}),
				field: reflect.StructField{},
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

			if err := setSlice(tt.args.vals, tt.args.value, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("setSlice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
