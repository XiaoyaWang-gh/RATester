package pagemeta

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewDefaultFrontmatterConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expected := FrontmatterConfig{
		Date:        []string{fmDate, fmPubDate, fmLastmod},
		Lastmod:     []string{fmGitAuthorDate, fmLastmod, fmDate, fmPubDate},
		PublishDate: []string{fmPubDate, fmDate},
		ExpiryDate:  []string{fmExpiryDate},
	}
	result := newDefaultFrontmatterConfig()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
