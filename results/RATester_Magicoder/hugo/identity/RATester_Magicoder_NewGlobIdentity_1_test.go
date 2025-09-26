package identity

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewGlobIdentity_1(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
		want Identity
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

			if got := NewGlobIdentity(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGlobIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}
