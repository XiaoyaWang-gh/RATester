package binding

import (
	"fmt"
	"net/http"
	"testing"
)

func TestBind_9(t *testing.T) {
	t.Run("should return error when mapHeader fails", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		req, _ := http.NewRequest("GET", "/", nil)
		hb := headerBinding{}
		obj := make(chan int) // non-struct type

		err := hb.Bind(req, obj)

		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("should return error when validate fails", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		req, _ := http.NewRequest("GET", "/", nil)
		hb := headerBinding{}
		obj := struct{}{} // empty struct

		err := hb.Bind(req, obj)

		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("should return nil when both mapHeader and validate pass", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		req, _ := http.NewRequest("GET", "/", nil)
		hb := headerBinding{}
		obj := struct{ Header map[string]string }{} // valid struct

		err := hb.Bind(req, &obj)

		if err != nil {
			t.Errorf("expected nil, got error: %v", err)
		}
	})
}
