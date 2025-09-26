package binding

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestBind_12(t *testing.T) {
	b := jsonBinding{}

	t.Run("should return error when request is nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		err := b.Bind(nil, nil)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("should return error when request body is nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		req, _ := http.NewRequest("GET", "/", nil)
		err := b.Bind(req, nil)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("should return no error when request is valid", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		req, _ := http.NewRequest("GET", "/", strings.NewReader("{}"))
		err := b.Bind(req, &struct{}{})
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})
}
