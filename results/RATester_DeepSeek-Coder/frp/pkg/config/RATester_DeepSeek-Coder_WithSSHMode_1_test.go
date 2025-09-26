package config

import (
	"fmt"
	"testing"
)

func TestWithSSHMode_1(t *testing.T) {
	type args struct {
		o *registerFlagOptions
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestWithSSHMode",
			args: args{
				o: &registerFlagOptions{
					sshMode: false,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			WithSSHMode()(tt.args.o)
			if got := tt.args.o.sshMode; got != tt.want {
				t.Errorf("WithSSHMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
