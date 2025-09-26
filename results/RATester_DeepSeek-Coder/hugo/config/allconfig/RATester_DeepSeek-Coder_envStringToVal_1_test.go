package allconfig

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEnvStringToVal_1(t *testing.T) {
	l := &configLoader{}

	t.Run("Test envStringToVal with 'disablekinds' key", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := l.envStringToVal("disablekinds", "kind1,kind2,kind3")
		if !reflect.DeepEqual(result, []string{"kind1", "kind2", "kind3"}) {
			t.Errorf("Expected result to be []string{\"kind1\", \"kind2\", \"kind3\"}, but got %v", result)
		}
	})

	t.Run("Test envStringToVal with 'disablelanguages' key", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := l.envStringToVal("disablelanguages", "lang1 lang2 lang3")
		if !reflect.DeepEqual(result, []string{"lang1", "lang2", "lang3"}) {
			t.Errorf("Expected result to be []string{\"lang1\", \"lang2\", \"lang3\"}, but got %v", result)
		}
	})

	t.Run("Test envStringToVal with default key", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := l.envStringToVal("default", "value")
		if result != "value" {
			t.Errorf("Expected result to be \"value\", but got %v", result)
		}
	})
}
