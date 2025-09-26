package hugo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNext_1(t *testing.T) {
	tests := []struct {
		name string
		v    Version
		want Version
	}{
		{
			name: "increments minor version",
			v:    Version{Major: 1, Minor: 2},
			want: Version{Major: 1, Minor: 3},
		},
		{
			name: "increments minor version",
			v:    Version{Major: 1, Minor: 2, PatchLevel: 3},
			want: Version{Major: 1, Minor: 3, PatchLevel: 3},
		},
		{
			name: "increments minor version",
			v:    Version{Major: 1, Minor: 2, PatchLevel: 3, Suffix: "beta"},
			want: Version{Major: 1, Minor: 3, PatchLevel: 3, Suffix: "beta"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.v.Next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
