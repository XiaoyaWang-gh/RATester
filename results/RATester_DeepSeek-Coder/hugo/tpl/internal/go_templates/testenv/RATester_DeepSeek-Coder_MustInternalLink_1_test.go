package testenv

import (
	"fmt"
	"testing"
)

func TestMustInternalLink_1(t *testing.T) {
	type args struct {
		t       testing.TB
		withCgo bool
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

			MustInternalLink(tt.args.t, tt.args.withCgo)
		})
	}
}
