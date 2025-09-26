package constraints

import (
	"fmt"
	"testing"
)

func TestLabelRegexFn_1(t *testing.T) {
	type args struct {
		name string
		expr string
	}
	tests := []struct {
		name string
		args args
		want constraintLabelFunc
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

			got := labelRegexFn(tt.args.name, tt.args.expr)
			if got == nil {
				t.Errorf("labelRegexFn() = %v, want %v", got, tt.want)
			}
		})
	}
}
