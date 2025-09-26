package identity

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewManager_1(t *testing.T) {
	type args struct {
		name string
		opts []ManagerOption
	}
	tests := []struct {
		name string
		args args
		want Manager
	}{
		{
			name: "test1",
			args: args{
				name: "test1",
				opts: []ManagerOption{
					func(m *identityManager) {
						m.onAddIdentity = func(id Identity) {
							// do something
						}
					},
				},
			},
			want: &identityManager{
				Identity: Anonymous,
				name:     "test1",
				ids:      Identities{},
				onAddIdentity: func(id Identity) {
					// do something
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

			if got := NewManager(tt.args.name, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewManager() = %v, want %v", got, tt.want)
			}
		})
	}
}
