package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEq_3(t *testing.T) {
	type args struct {
		arg1 reflect.Value
		arg2 []reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				arg1: reflect.ValueOf(1),
				arg2: []reflect.Value{reflect.ValueOf(1)},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "case2",
			args: args{
				arg1: reflect.ValueOf(1),
				arg2: []reflect.Value{reflect.ValueOf(2)},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "case3",
			args: args{
				arg1: reflect.ValueOf(1),
				arg2: []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)},
			},
			want:    false,
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

			got, err := eq(tt.args.arg1, tt.args.arg2...)
			if (err != nil) != tt.wantErr {
				t.Errorf("eq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("eq() = %v, want %v", got, tt.want)
			}
		})
	}
}
