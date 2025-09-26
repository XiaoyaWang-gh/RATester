package htesting

import (
	"fmt"
	"os"
	"testing"

	"github.com/alecthomas/assert"
)

func TestSupportsAll_1(t *testing.T) {
	t.Run("IsGitHubAction", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Run("true", func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			os.Setenv("GITHUB_ACTIONS", "true")
			defer os.Unsetenv("GITHUB_ACTIONS")
			assert.True(t, SupportsAll())
		})
		t.Run("false", func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			os.Setenv("GITHUB_ACTIONS", "false")
			defer os.Unsetenv("GITHUB_ACTIONS")
			assert.False(t, SupportsAll())
		})
	})
	t.Run("CI_LOCAL", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Run("true", func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			os.Setenv("CI_LOCAL", "true")
			defer os.Unsetenv("CI_LOCAL")
			assert.True(t, SupportsAll())
		})
		t.Run("false", func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			os.Setenv("CI_LOCAL", "false")
			defer os.Unsetenv("CI_LOCAL")
			assert.False(t, SupportsAll())
		})
	})
}
