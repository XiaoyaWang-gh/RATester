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

	site := testSite{}
	config := site.Config()

	if config == (SiteConfig{}) {
		t.Errorf("Expected non-empty SiteConfig, got empty SiteConfig")
	}
}
