package gin

import (
	"fmt"
	"testing"
)

func TestStaticFileHandler_1(t *testing.T) {
	router := new(RouterGroup)
	handler := func(c *Context) {}

	t.Run("Test staticFileHandler with valid relativePath", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("The code panicked and it should not have. %v", r)
			}
		}()
		router.staticFileHandler("validPath", handler)
	})

	t.Run("Test staticFileHandler with invalid relativePath containing ':'", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic as expected.")
			}
		}()
		router.staticFileHandler("invalid:Path", handler)
	})

	t.Run("Test staticFileHandler with invalid relativePath containing '*'", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic as expected.")
			}
		}()
		router.staticFileHandler("invalid*Path", handler)
	})
}
