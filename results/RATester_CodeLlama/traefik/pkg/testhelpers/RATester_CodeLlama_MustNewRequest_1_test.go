package testhelpers

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustNewRequest_1(t *testing.T) {
	t.Parallel()
	t.Run("method is empty", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()
		assert.PanicsWithValue(t, "failed to create HTTP GET Request for '': method is empty", func() {
			MustNewRequest("", "https://example.com", nil)
		})
	})
	t.Run("urlStr is empty", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()
		assert.PanicsWithValue(t, "failed to create HTTP GET Request for '': urlStr is empty", func() {
			MustNewRequest("GET", "", nil)
		})
	})
	t.Run("body is nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()
		assert.NotPanics(t, func() {
			MustNewRequest("GET", "https://example.com", nil)
		})
	})
	t.Run("body is not nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()
		assert.NotPanics(t, func() {
			MustNewRequest("GET", "https://example.com", strings.NewReader(""))
		})
	})
}
