package partials

import (
	"fmt"
	"testing"
)

func TestTemplateName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	k := partialCacheKey{Name: "test"}
	if k.templateName() != "partials/test" {
		t.Errorf("templateName() = %v, want %v", k.templateName(), "partials/test")
	}
}
