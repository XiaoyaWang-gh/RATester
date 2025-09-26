package legacy

import (
	"fmt"
	"testing"

	"gopkg.in/ini.v1"
)

func TestCopySection_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	source := ini.Empty()
	target := ini.Empty()
	source.NewSection("section")
	source.Section("section").NewKey("key", "value")
	copySection(source.Section("section"), target.Section("section"))
	if target.Section("section").Key("key").String() != "value" {
		t.Errorf("copySection failed")
	}
}
