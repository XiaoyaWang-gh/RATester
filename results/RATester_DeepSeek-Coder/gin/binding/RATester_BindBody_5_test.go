package binding

import (
	"fmt"
	"testing"
)

func TestBindBody_5(t *testing.T) {
	b := jsonBinding{}
	type testStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		body := []byte(`{"name":"John", "age":30}`)
		obj := &testStruct{}
		err := b.BindBody(body, obj)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if obj.Name != "John" || obj.Age != 30 {
			t.Errorf("expected Name to be 'John' and Age to be 30, got %v", obj)
		}
	})

	t.Run("failure", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		body := []byte(`{"name":"John", "age":"thirty"}`)
		obj := &testStruct{}
		err := b.BindBody(body, obj)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
