package legacy

import (
	"fmt"
	"testing"

	"gopkg.in/ini.v1"
)

func TestRenderRangeProxyTemplates_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := ini.Empty()
	section := f.Section("range:test")
	section.Key("local_port").SetValue("1000-1005")
	section.Key("remote_port").SetValue("2000-2005")
	err := renderRangeProxyTemplates(f, section)
	if err != nil {
		t.Errorf("renderRangeProxyTemplates error: %v", err)
	}
	if len(f.Sections()) != 6 {
		t.Errorf("renderRangeProxyTemplates error: %v", f.Sections())
	}
}
