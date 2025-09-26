package gin

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestJSON_3(t *testing.T) {
	t.Run("should return nil when errorMsgs is empty", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var errs errorMsgs
		result := errs.JSON()
		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})

	t.Run("should return JSON representation of the last error when errorMsgs has one error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		errs := errorMsgs{&Error{Err: errors.New("test error"), Type: 1, Meta: "meta"}}
		result := errs.JSON()
		expected := map[string]any{"Err": "test error", "Type": 1, "Meta": "meta"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("should return JSON representation of all errors when errorMsgs has multiple errors", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		errs := errorMsgs{
			&Error{Err: errors.New("test error 1"), Type: 1, Meta: "meta 1"},
			&Error{Err: errors.New("test error 2"), Type: 2, Meta: "meta 2"},
		}
		result := errs.JSON()
		expected := []any{
			map[string]any{"Err": "test error 1", "Type": 1, "Meta": "meta 1"},
			map[string]any{"Err": "test error 2", "Type": 2, "Meta": "meta 2"},
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
