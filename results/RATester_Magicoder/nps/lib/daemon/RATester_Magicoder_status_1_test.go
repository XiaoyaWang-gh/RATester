package daemon

import (
	"fmt"
	"testing"
)

func Teststatus_1(t *testing.T) {
	type args struct {
		f       string
		pidPath string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := status(tt.args.f, tt.args.pidPath); got != tt.want {
				t.Errorf("status() = %v, want %v", got, tt.want)
			}
		})
	}
}
