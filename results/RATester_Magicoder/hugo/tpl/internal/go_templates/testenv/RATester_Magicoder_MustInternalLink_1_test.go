package testenv

import (
	"fmt"
	"testing"
)

func TestMustInternalLink_1(t *testing.T) {
	type args struct {
		withCgo bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{withCgo: true},
		},
		{
			name: "Test case 2",
			args: args{withCgo: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			MustInternalLink(t, tt.args.withCgo)
		})
	}
}
