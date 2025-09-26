package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGe_1(t *testing.T) {
	type args struct {
		arg1 reflect.Value
		arg2 reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				arg1: reflect.ValueOf(1),
				arg2: reflect.ValueOf(2),
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				arg1: reflect.ValueOf(2),
				arg2: reflect.ValueOf(1),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "test case 3",
			args: args{
				arg1: reflect.ValueOf(1),
				arg2: reflect.ValueOf(1),
			},
			want:    false,
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

			got, err := ge(tt.args.arg1, tt.args.arg2)
			if (err != nil) != tt.wantErr {
				t.Errorf("ge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ge() = %v, want %v", got, tt.want)
			}
		})
	}
}
