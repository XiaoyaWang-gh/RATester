package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIndexArg_1(t *testing.T) {
	type args struct {
		index reflect.Value
		cap   int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				index: reflect.ValueOf(1),
				cap:   10,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "case2",
			args: args{
				index: reflect.ValueOf(10),
				cap:   10,
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "case3",
			args: args{
				index: reflect.ValueOf(10),
				cap:   1,
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

			got, err := indexArg(tt.args.index, tt.args.cap)
			if (err != nil) != tt.wantErr {
				t.Errorf("indexArg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("indexArg() = %v, want %v", got, tt.want)
			}
		})
	}
}
