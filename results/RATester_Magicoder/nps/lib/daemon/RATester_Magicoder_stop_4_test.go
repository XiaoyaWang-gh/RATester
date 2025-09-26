package daemon

import (
	"fmt"
	"testing"
)

func Teststop_4(t *testing.T) {
	type args struct {
		f       string
		p       string
		pidPath string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Teststop_0",
			args: args{
				f:       "test",
				p:       "test.exe",
				pidPath: "test",
			},
		},
		{
			name: "Teststop_1",
			args: args{
				f:       "test",
				p:       "test.exe",
				pidPath: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			stop(tt.args.f, tt.args.p, tt.args.pidPath)
		})
	}
}
