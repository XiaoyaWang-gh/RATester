package allconfig

import (
	"fmt"
	"testing"
)

func TestTemplateMetricsHints_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	if c.TemplateMetricsHints() {
		t.Error("TemplateMetricsHints() should be false")
	}
}
