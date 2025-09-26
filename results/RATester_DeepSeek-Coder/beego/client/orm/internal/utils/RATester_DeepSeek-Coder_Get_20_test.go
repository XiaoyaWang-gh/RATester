package utils

import (
	"fmt"
	"testing"
)

func TestGet_20(t *testing.T) {
	type fields struct {
		Arg ArgInt
	}
	type args struct {
		i    int
		args []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantR  int
	}{
		{
			name: "TestGet_0",
			fields: fields{
				Arg: ArgInt{1, 2, 3, 4, 5},
			},
			args: args{
				i: 2,
			},
			wantR: 3,
		},
		{
			name: "TestGet_1",
			fields: fields{
				Arg: ArgInt{1, 2, 3, 4, 5},
			},
			args: args{
				i:    2,
				args: []int{6},
			},
			wantR: 6,
		},
		{
			name: "TestGet_2",
			fields: fields{
				Arg: ArgInt{1, 2, 3, 4, 5},
			},
			args: args{
				i: -1,
			},
			wantR: 0,
		},
		{
			name: "TestGet_3",
			fields: fields{
				Arg: ArgInt{1, 2, 3, 4, 5},
			},
			args: args{
				i: 5,
			},
			wantR: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			a := ArgInt(tt.fields.Arg)
			if gotR := a.Get(tt.args.i, tt.args.args...); gotR != tt.wantR {
				t.Errorf("Get() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
