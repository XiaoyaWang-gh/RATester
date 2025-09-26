package allconfig

import (
	"fmt"
	"testing"
)

func TestTemplateMetrics_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	if c.TemplateMetrics() {
		t.Error("TemplateMetrics() should be false")
	}
}
