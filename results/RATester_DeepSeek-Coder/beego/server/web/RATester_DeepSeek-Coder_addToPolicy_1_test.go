package web

import (
	"fmt"
	"testing"
)

func TestAddToPolicy_1(t *testing.T) {
	type args struct {
		method  string
		pattern string
		r       []PolicyFunc
	}
	tests := []struct {
		name string
		args args
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

			p := &ControllerRegister{}
			p.addToPolicy(tt.args.method, tt.args.pattern, tt.args.r...)
		})
	}
}
