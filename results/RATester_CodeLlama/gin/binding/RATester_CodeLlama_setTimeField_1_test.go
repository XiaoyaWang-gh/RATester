package binding

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestSetTimeField_1(t *testing.T) {
	type args struct {
		val         string
		structField reflect.StructField
		value       reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				val:         "2020-01-01T12:00:00Z",
				structField: reflect.StructField{},
				value:       reflect.ValueOf(time.Time{}),
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				val:         "2020-01-01T12:00:00Z",
				structField: reflect.StructField{},
				value:       reflect.ValueOf(time.Time{}),
			},
			wantErr: false,
		},
		{
			name: "test case 3",
			args: args{
				val:         "2020-01-01T12:00:00Z",
				structField: reflect.StructField{},
				value:       reflect.ValueOf(time.Time{}),
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

			if err := setTimeField(tt.args.val, tt.args.structField, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("setTimeField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
