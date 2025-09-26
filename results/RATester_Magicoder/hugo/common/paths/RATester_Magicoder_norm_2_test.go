package paths

import (
	"fmt"
	"testing"
)

func Testnorm_2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		p    *Path
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

			if got := tt.p.norm(tt.args.s); got != tt.want {
				t.Errorf("norm() = %v, want %v", got, tt.want)
			}
		})
	}
}
