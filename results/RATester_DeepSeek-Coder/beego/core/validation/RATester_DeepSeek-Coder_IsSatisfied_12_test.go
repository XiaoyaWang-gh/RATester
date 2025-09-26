package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_12(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		p    Phone
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

			if got := tt.p.IsSatisfied(tt.args.obj); got != tt.want {
				t.Errorf("Phone.IsSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}
