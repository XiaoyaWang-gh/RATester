package hugo

import (
	"fmt"
	"testing"
)

func TestGetBuildInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bi := getBuildInfo()
	if bi == nil {
		t.Error("getBuildInfo() returned nil")
	}
	if bi.BuildInfo == nil {
		t.Error("getBuildInfo().BuildInfo is nil")
	}
	if bi.VersionControlSystem == "" {
		t.Error("getBuildInfo().VersionControlSystem is empty")
	}
	if bi.Revision == "" {
		t.Error("getBuildInfo().Revision is empty")
	}
	if bi.RevisionTime == "" {
		t.Error("getBuildInfo().RevisionTime is empty")
	}
	if bi.Modified == false {
		t.Error("getBuildInfo().Modified is false")
	}
	if bi.GoOS == "" {
		t.Error("getBuildInfo().GoOS is empty")
	}
	if bi.GoArch == "" {
		t.Error("getBuildInfo().GoArch is empty")
	}
}
