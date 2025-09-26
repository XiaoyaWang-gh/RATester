package resources

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/identity"
)

func TestForEeachIdentityByName_1(t *testing.T) {
	type args struct {
		name string
		f    func(identity.Identity) bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				name: "test",
				f: func(identity.Identity) bool {
					return true
				},
			},
		},
		{
			name: "Test case 2",
			args: args{
				name: "test2",
				f: func(identity.Identity) bool {
					return false
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := &resourceAdapter{}
			r.ForEeachIdentityByName(tt.args.name, tt.args.f)
		})
	}
}
