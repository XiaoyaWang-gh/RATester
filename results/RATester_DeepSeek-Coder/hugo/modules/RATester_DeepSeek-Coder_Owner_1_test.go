package modules

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOwner_1(t *testing.T) {
	type fields struct {
		owner Module
	}
	tests := []struct {
		name   string
		fields fields
		want   Module
	}{
		{
			name: "Test case 1",
			fields: fields{
				owner: &moduleAdapter{
					path:    "path",
					dir:     "dir",
					version: "version",
				},
			},
			want: &moduleAdapter{
				path:    "path",
				dir:     "dir",
				version: "version",
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

			m := &moduleAdapter{
				owner: tt.fields.owner,
			}
			if got := m.Owner(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moduleAdapter.Owner() = %v, want %v", got, tt.want)
			}
		})
	}
}
