package page

import (
	"fmt"
	"testing"
)

func TestDisqusShortname_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	site := testSite{}
	shortname := site.DisqusShortname()
	if shortname != "" {
		t.Errorf("Expected empty string, but got %s", shortname)
	}
}
