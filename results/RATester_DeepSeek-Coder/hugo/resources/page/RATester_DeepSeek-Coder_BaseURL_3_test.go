package page

import (
	"fmt"
	"testing"
)

func TestBaseURL_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	site := testSite{}
	expected := ""
	actual := site.BaseURL()
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
