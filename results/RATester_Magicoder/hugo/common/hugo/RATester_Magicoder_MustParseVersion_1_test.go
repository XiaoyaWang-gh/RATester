package hugo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMustParseVersion_1(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    Version
	}{
		{
			name:    "valid version",
			version: "1.2.3",
			want:    Version{Major: 1, Minor: 2, PatchLevel: 3},
		},
		{
			name:    "invalid version",
			version: "invalid",
			want:    Version{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := MustParseVersion(tt.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustParseVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
