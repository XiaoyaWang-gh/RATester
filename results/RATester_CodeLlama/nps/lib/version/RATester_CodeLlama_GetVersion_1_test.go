package version

import (
	"fmt"
	"testing"
)

func TestGetVersion_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	version := GetVersion()
	if version != "0.26.0" {
		t.Errorf("GetVersion() = %v, want %v", version, "0.26.0")
	}
}
