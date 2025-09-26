package web

import (
	"fmt"
	"testing"
)

func TestLe_1(t *testing.T) {
	type args struct {
		arg1 interface{}
		arg2 interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				arg1: 1,
				arg2: 2,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				arg1: 2,
				arg2: 1,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				arg1: 1,
				arg2: 1,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Test Case 4",
			args: args{
				arg1: "1",
				arg2: 1,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Test Case 5",
			args: args{
				arg1: 1,
				arg2: "1",
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

			got, err := le(tt.args.arg1, tt.args.arg2)
			if (err != nil) != tt.wantErr {
				t.Errorf("le() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("le() = %v, want %v", got, tt.want)
			}
		})
	}
}
