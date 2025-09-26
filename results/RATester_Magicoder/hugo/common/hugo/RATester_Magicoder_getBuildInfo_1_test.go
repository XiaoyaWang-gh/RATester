package hugo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestgetBuildInfo_1(t *testing.T) {
	tests := []struct {
		name string
		want *buildInfo
	}{
		{
			name: "Test Case 1",
			want: &buildInfo{
				VersionControlSystem: "git",
				Revision:             "abc123",
				RevisionTime:         "2022-01-01T00:00:00Z",
				Modified:             false,
				GoOS:                 "linux",
				GoArch:               "amd64",
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

			if got := getBuildInfo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBuildInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
