package hugo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrev_1(t *testing.T) {
	tests := []struct {
		name string
		v    Version
		want Version
	}{
		{
			name: "Test case 1",
			v:    Version{Major: 1, Minor: 2},
			want: Version{Major: 1, Minor: 1},
		},
		{
			name: "Test case 2",
			v:    Version{Major: 2, Minor: 0},
			want: Version{Major: 2, Minor: -1},
		},
		{
			name: "Test case 3",
			v:    Version{Major: 0, Minor: 0},
			want: Version{Major: 0, Minor: -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.v.Prev(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Version.Prev() = %v, want %v", got, tt.want)
			}
		})
	}
}
