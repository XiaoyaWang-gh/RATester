package maps

import (
	"fmt"
	"testing"
)

func TestGetNewKey_1(t *testing.T) {
	type args struct {
		keyPath string
	}
	tests := []struct {
		name string
		r    KeyRenamer
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

			if got := tt.r.getNewKey(tt.args.keyPath); got != tt.want {
				t.Errorf("getNewKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
