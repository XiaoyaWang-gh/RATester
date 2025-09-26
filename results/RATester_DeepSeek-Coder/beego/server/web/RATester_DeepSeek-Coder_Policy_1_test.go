package web

import (
	"fmt"
	"testing"
)

func TestPolicy_1(t *testing.T) {
	type args struct {
		pattern string
		method  string
		policy  []PolicyFunc
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

			Policy(tt.args.pattern, tt.args.method, tt.args.policy...)
		})
	}
}
