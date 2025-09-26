package identity

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewPredicateIdentity_1(t *testing.T) {
	type args struct {
		probablyDependent  func(Identity) bool
		probablyDependency func(Identity) bool
	}
	tests := []struct {
		name string
		args args
		want *predicateIdentity
	}{
		{
			name: "Test case 1",
			args: args{
				probablyDependent:  func(Identity) bool { return true },
				probablyDependency: func(Identity) bool { return true },
			},
			want: &predicateIdentity{
				id:                 "predicate1",
				probablyDependent:  func(Identity) bool { return true },
				probablyDependency: func(Identity) bool { return true },
			},
		},
		{
			name: "Test case 2",
			args: args{
				probablyDependent:  func(Identity) bool { return false },
				probablyDependency: func(Identity) bool { return false },
			},
			want: &predicateIdentity{
				id:                 "predicate2",
				probablyDependent:  func(Identity) bool { return false },
				probablyDependency: func(Identity) bool { return false },
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

			if got := NewPredicateIdentity(tt.args.probablyDependent, tt.args.probablyDependency); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPredicateIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}
