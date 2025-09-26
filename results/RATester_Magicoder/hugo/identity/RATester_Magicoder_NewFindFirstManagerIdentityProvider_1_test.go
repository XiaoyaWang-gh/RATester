package identity

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewFindFirstManagerIdentityProvider_1(t *testing.T) {
	type args struct {
		m  Manager
		id Identity
	}
	tests := []struct {
		name string
		args args
		want FindFirstManagerIdentityProvider
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

			if got := NewFindFirstManagerIdentityProvider(tt.args.m, tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFindFirstManagerIdentityProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}
