package page

import (
	"fmt"
	"testing"
)

func TestNewPermalinkExpander_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	urlize := func(uri string) string {
		return uri
	}

	patterns := map[string]map[string]string{
		"section": {
			"permalink": "/:year/:month/:day/:slug",
		},
	}

	_, err := NewPermalinkExpander(urlize, patterns)
	if err != nil {
		t.Errorf("NewPermalinkExpander() error = %v", err)
		return
	}
}
