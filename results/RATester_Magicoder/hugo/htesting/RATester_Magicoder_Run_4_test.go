package htesting

import (
	"fmt"
	"regexp"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestRun_4(t *testing.T) {
	r := &PinnedRunner{
		c:  &qt.C{TB: t},
		re: regexp.MustCompile("pinned"),
	}

	t.Run("pinned test", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		r.Run("pinned test", func(c *qt.C) {
			// test code here
		})
	})

	t.Run("non-pinned test", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		r.Run("non-pinned test", func(c *qt.C) {
			// test code here
		})
	})
}
