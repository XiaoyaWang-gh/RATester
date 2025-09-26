package paths

import (
	"fmt"
	"testing"
)

func TestReset_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{}
	p.reset()
	if p.s != "" {
		t.Errorf("p.s = %q, want %q", p.s, "")
	}
	if p.posContainerLow != -1 {
		t.Errorf("p.posContainerLow = %d, want %d", p.posContainerLow, -1)
	}
	if p.posContainerHigh != -1 {
		t.Errorf("p.posContainerHigh = %d, want %d", p.posContainerHigh, -1)
	}
	if p.posSectionHigh != -1 {
		t.Errorf("p.posSectionHigh = %d, want %d", p.posSectionHigh, -1)
	}
	if p.component != "" {
		t.Errorf("p.component = %q, want %q", p.component, "")
	}
	if p.bundleType != 0 {
		t.Errorf("p.bundleType = %d, want %d", p.bundleType, 0)
	}
	if len(p.identifiers) != 0 {
		t.Errorf("len(p.identifiers) = %d, want %d", len(p.identifiers), 0)
	}
	if p.posIdentifierLanguage != -1 {
		t.Errorf("p.posIdentifierLanguage = %d, want %d", p.posIdentifierLanguage, -1)
	}
	if p.disabled != false {
		t.Errorf("p.disabled = %t, want %t", p.disabled, false)
	}
	if p.trimLeadingSlash != false {
		t.Errorf("p.trimLeadingSlash = %t, want %t", p.trimLeadingSlash, false)
	}
	if p.unnormalized != nil {
		t.Errorf("p.unnormalized = %v, want %v", p.unnormalized, nil)
	}
}
