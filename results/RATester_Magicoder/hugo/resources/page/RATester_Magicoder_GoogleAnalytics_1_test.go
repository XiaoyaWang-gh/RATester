package page

import (
	"fmt"
	"testing"
)

func TestGoogleAnalytics_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	site := testSite{}
	result := site.GoogleAnalytics()
	if result != "" {
		t.Errorf("Expected empty string, but got %s", result)
	}
}
