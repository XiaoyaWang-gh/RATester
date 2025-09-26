package resource

import (
	"fmt"
	"testing"
)

func TestNameNormalizedOrName_1(t *testing.T) {
	type args struct {
		r Resource
	}
	tests := []struct {
		name string
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

			if got := NameNormalizedOrName(tt.args.r); got != tt.want {
				t.Errorf("NameNormalizedOrName() = %v, want %v", got, tt.want)
			}
		})
	}
}
