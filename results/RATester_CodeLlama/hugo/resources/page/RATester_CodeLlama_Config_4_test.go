package page

import (
	"fmt"
	"testing"
)

func TestConfig_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := testSite{}
	if s.Config() != (SiteConfig{}) {
		t.Errorf("Config() = %v, want %v", s.Config(), SiteConfig{})
	}
}
