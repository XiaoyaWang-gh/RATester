package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCall_1(t *testing.T) {
	type args struct {
		name string
		fn   reflect.Value
		args []reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		want    reflect.Value
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				name: "case1",
				fn:   reflect.ValueOf(func(a int, b int) int { return a + b }),
				args: []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)},
			},
			want:    reflect.ValueOf(3),
			wantErr: false,
		},
		{
			name: "case2",
			args: args{
				name: "case2",
				fn:   reflect.ValueOf(func(a int, b int) int { return a + b }),
				args: []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2), reflect.ValueOf(3)},
			},
			want:    reflect.ValueOf(3),
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

			got, err := call(tt.args.name, tt.args.fn, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("call() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("call() got = %v, want %v", got, tt.want)
			}
		})
	}
}
