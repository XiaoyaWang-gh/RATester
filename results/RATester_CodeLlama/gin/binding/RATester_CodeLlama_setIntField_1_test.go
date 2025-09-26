package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetIntField_1(t *testing.T) {
	type args struct {
		val        string
		bitSize    int
		fieldValue reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				val:        "1",
				bitSize:    10,
				fieldValue: reflect.ValueOf(1),
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				val:        "1",
				bitSize:    10,
				fieldValue: reflect.ValueOf(1),
			},
			wantErr: false,
		},
		{
			name: "test case 3",
			args: args{
				val:        "1",
				bitSize:    10,
				fieldValue: reflect.ValueOf(1),
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

			if err := setIntField(tt.args.val, tt.args.bitSize, tt.args.fieldValue); (err != nil) != tt.wantErr {
				t.Errorf("setIntField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
