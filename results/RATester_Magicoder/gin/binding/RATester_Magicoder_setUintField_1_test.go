package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestsetUintField_1(t *testing.T) {
	type args struct {
		val     string
		bitSize int
		field   reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				val:     "10",
				bitSize: 64,
				field:   reflect.ValueOf(uint64(0)),
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				val:     "",
				bitSize: 64,
				field:   reflect.ValueOf(uint64(0)),
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				val:     "invalid",
				bitSize: 64,
				field:   reflect.ValueOf(uint64(0)),
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

			if err := setUintField(tt.args.val, tt.args.bitSize, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("setUintField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
