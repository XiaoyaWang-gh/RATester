package fmt

import (
	"fmt"
	"testing"
)

func TestErrorf_1(t *testing.T) {
	type args struct {
		format string
		args   []any
	}
	tests := []struct {
		name string
		ns   *Namespace
		args args
		want string
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

			if got := tt.ns.Errorf(tt.args.format, tt.args.args...); got != tt.want {
				t.Errorf("Errorf() = %v, want %v", got, tt.want)
			}
		})
	}
}
