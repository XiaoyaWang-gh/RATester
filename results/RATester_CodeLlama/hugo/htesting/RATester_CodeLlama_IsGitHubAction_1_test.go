package htesting

import (
	"fmt"
	"os"
	"testing"

	"github.com/alecthomas/assert"
)

func TestIsGitHubAction_1(t *testing.T) {
	t.Run("GITHUB_ACTION is set", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("GITHUB_ACTION", "true")
		defer os.Unsetenv("GITHUB_ACTION")
		assert.True(t, IsGitHubAction())
	})
	t.Run("GITHUB_ACTION is not set", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Unsetenv("GITHUB_ACTION")
		assert.False(t, IsGitHubAction())
	})
}
