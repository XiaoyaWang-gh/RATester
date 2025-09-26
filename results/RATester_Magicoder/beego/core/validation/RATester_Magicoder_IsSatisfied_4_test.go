package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_4(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		n    NoMatch
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

			if got := tt.n.IsSatisfied(tt.args.obj); got != tt.want {
				t.Errorf("IsSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}
