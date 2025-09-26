package httplib

import (
	"fmt"
	"testing"
)

func TestBody_1(t *testing.T) {
	b := &BeegoHTTPRequest{}

	t.Run("TestBody_String", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b.Body("test")
		if b.req.Body == nil {
			t.Errorf("Expected request body to be set, but it was not")
		}
	})

	t.Run("TestBody_Bytes", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b.Body([]byte("test"))
		if b.req.Body == nil {
			t.Errorf("Expected request body to be set, but it was not")
		}
	})

	t.Run("TestBody_UnsupportedType", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected a panic when using an unsupported type")
			}
		}()
		b.Body(123)
	})
}
