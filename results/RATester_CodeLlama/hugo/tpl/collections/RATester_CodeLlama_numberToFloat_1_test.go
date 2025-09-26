package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNumberToFloat_1(t *testing.T) {
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				v: reflect.ValueOf(1),
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "case 2",
			args: args{
				v: reflect.ValueOf(1.1),
			},
			want:    1.1,
			wantErr: false,
		},
		{
			name: "case 3",
			args: args{
				v: reflect.ValueOf("1.1"),
			},
			want:    1.1,
			wantErr: false,
		},
		{
			name: "case 4",
			args: args{
				v: reflect.ValueOf([]int{1, 2, 3}),
			},
			want:    0,
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

			got, err := numberToFloat(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("numberToFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("numberToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
