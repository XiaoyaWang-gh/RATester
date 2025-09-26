package hugo

import (
	"fmt"
	"testing"
)

func TestGetBuildInfo_1(t *testing.T) {
	tests := []struct {
		name string
		want *buildInfo
	}{
		{
			name: "Test case 1",
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

			got := getBuildInfo()
			if got.VersionControlSystem != tt.want.VersionControlSystem {
				t.Errorf("getBuildInfo() = %v, want %v", got.VersionControlSystem, tt.want.VersionControlSystem)
			}
			if got.Revision != tt.want.Revision {
				t.Errorf("getBuildInfo() = %v, want %v", got.Revision, tt.want.Revision)
			}
			if got.RevisionTime != tt.want.RevisionTime {
				t.Errorf("getBuildInfo() = %v, want %v", got.RevisionTime, tt.want.RevisionTime)
			}
			if got.Modified != tt.want.Modified {
				t.Errorf("getBuildInfo() = %v, want %v", got.Modified, tt.want.Modified)
			}
			if got.GoOS != tt.want.GoOS {
				t.Errorf("getBuildInfo() = %v, want %v", got.GoOS, tt.want.GoOS)
			}
			if got.GoArch != tt.want.GoArch {
				t.Errorf("getBuildInfo() = %v, want %v", got.GoArch, tt.want.GoArch)
			}
		})
	}
}
