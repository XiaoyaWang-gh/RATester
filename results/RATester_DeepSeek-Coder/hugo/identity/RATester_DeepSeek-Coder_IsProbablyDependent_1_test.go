package identity

import (
	"fmt"
	"testing"
)

func TestIsProbablyDependent_1(t *testing.T) {
	type args struct {
		other Identity
	}
	tests := []struct {
		name string
		id   *predicateIdentity
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

			if got := tt.id.IsProbablyDependent(tt.args.other); got != tt.want {
				t.Errorf("predicateIdentity.IsProbablyDependent() = %v, want %v", got, tt.want)
			}
		})
	}
}
