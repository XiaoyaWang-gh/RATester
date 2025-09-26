package identity

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindFirstManagerIdentity_1(t *testing.T) {
	type fields struct {
		Identity        Identity
		ManagerIdentity ManagerIdentity
	}
	tests := []struct {
		name   string
		fields fields
		want   ManagerIdentity
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

			f := findFirstManagerIdentity{
				Identity:        tt.fields.Identity,
				ManagerIdentity: tt.fields.ManagerIdentity,
			}
			if got := f.FindFirstManagerIdentity(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFirstManagerIdentity.FindFirstManagerIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}
