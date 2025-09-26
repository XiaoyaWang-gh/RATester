package hugo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMustParseVersion_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want Version
	}{
		{
			name: "Test case 1",
			args: args{s: "0.104.0"},
			want: Version{Major: 0, Minor: 104, PatchLevel: 0},
		},
		{
			name: "Test case 2",
			args: args{s: "0.104.1"},
			want: Version{Major: 0, Minor: 104, PatchLevel: 1},
		},
		{
			name: "Test case 3",
			args: args{s: "0.104.2-DEV"},
			want: Version{Major: 0, Minor: 104, PatchLevel: 2, Suffix: "DEV"},
		},
		{
			name: "Test case 4",
			args: args{s: "0.104.2-DEV+abc"},
			want: Version{Major: 0, Minor: 104, PatchLevel: 2, Suffix: "DEV+abc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := MustParseVersion(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustParseVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
