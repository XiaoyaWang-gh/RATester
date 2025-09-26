package hugolib

import (
	"fmt"
	"testing"
)

func TestEditFileReplaceAll_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := NewIntegrationTestBuilder(IntegrationTestConfig{
		T: t,
		TxtarString: `
-- test.txt --
Hello, World!
`,
	})

	b.EditFileReplaceAll("test.txt", "Hello", "Hi")

	b.AssertFileContent("test.txt", "Hi, World!")
}
