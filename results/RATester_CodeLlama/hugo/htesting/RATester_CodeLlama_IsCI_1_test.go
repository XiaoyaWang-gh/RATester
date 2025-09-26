package htesting

import (
	"fmt"
	"os"
	"testing"

	"github.com/alecthomas/assert"
)

func TestIsCI_1(t *testing.T) {
	t.Run("CI", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("CI", "true")
		defer os.Unsetenv("CI")
		assert.True(t, IsCI())
	})

	t.Run("CI_LOCAL", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("CI_LOCAL", "true")
		defer os.Unsetenv("CI_LOCAL")
		assert.True(t, IsCI())
	})

	t.Run("CIRCLE_BRANCH", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("CIRCLE_BRANCH", "true")
		defer os.Unsetenv("CIRCLE_BRANCH")
		assert.False(t, IsCI())
	})
}
