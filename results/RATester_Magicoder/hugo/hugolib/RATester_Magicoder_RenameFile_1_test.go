package hugolib

import (
	"fmt"
	"testing"
)

func TestRenameFile_1(t *testing.T) {
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
	b.Init()
	b.RenameFile("test.txt", "new_test.txt")
	b.AssertFileContent("new_test.txt", "Hello, World!")
}
