package modules

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFilterUnwantedMounts_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mounts := []Mount{
		{
			Source: "scss",
			Target: "assets/bootstrap/scss",
		},
		{
			Source: "scss",
			Target: "assets/bootstrap/scss",
		},
		{
			Source: "scss",
			Target: "assets/bootstrap/scss",
		},
	}

	expected := []Mount{
		{
			Source: "scss",
			Target: "assets/bootstrap/scss",
		},
	}

	actual := filterUnwantedMounts(mounts)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("filterUnwantedMounts(%v) = %v, want %v", mounts, actual, expected)
	}
}
