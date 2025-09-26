package page

import (
	"fmt"
	"testing"
)

func TestBaseURL_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	site := testSite{}
	expected := ""
	result := site.BaseURL()
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
