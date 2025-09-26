package identity

import (
	"fmt"
	"testing"
)

func TestForEeachIdentity_1(t *testing.T) {
	type args struct {
		f func(id Identity) bool
	}
	tests := []struct {
		name string
		m    *nopManager
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

			if got := tt.m.forEeachIdentity(tt.args.f); got != tt.want {
				t.Errorf("nopManager.forEeachIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}
