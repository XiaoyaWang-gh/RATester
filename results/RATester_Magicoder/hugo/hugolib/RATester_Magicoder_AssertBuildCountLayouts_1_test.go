package hugolib

import (
	"fmt"
	"testing"
)

func TestAssertBuildCountLayouts_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := NewIntegrationTestBuilder(IntegrationTestConfig{
		T: t,
		TxtarString: `
-- test.txt --
hello world
-- config.toml --
[build]
disableKinds = ["RSS", "Sitemap", "Section"]
`,
	})

	b.AssertBuildCountLayouts(1)
}
