package identity

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithOnAddIdentity_1(t *testing.T) {
	type args struct {
		f func(id Identity)
	}
	tests := []struct {
		name string
		args args
		want ManagerOption
	}{
		{
			name: "Test case 1",
			args: args{
				f: func(id Identity) {
					// Test function implementation
				},
			},
			want: func(m *identityManager) {
				m.onAddIdentity = func(id Identity) {
					// Test function implementation
				}
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := WithOnAddIdentity(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithOnAddIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}
