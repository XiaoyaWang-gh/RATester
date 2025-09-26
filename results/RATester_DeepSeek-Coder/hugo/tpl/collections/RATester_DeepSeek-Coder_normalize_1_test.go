package collections

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/hashing"
	"github.com/gohugoio/hugo/common/types"
)

func TestNormalize_1(t *testing.T) {
	t.Run("Test with non-comparable type", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testValue := reflect.ValueOf(struct{}{})
		expected := hashing.HashUint64(testValue.Interface())
		result := normalize(testValue)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Test with number type", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testValue := reflect.ValueOf(123)
		expected, _ := numberToFloat(testValue)
		result := normalize(testValue)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Test with non-number type", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testValue := reflect.ValueOf("test")
		expected := types.Unwrapv(testValue.Interface())
		result := normalize(testValue)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
