package images

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
)

func TestToFilters_1(t *testing.T) {
	t.Run("Testing with []gift.Filter", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		filters := []gift.Filter{
			gift.Contrast(2),
			gift.Gamma(2),
		}
		result := ToFilters(filters)
		if !reflect.DeepEqual(result, filters) {
			t.Errorf("Expected %v, got %v", filters, result)
		}
	})

	t.Run("Testing with []filter", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		filters := []filter{
			{Options: filterOpts{Version: 1, Vals: 2}, Filter: gift.Contrast(2)},
			{Options: filterOpts{Version: 1, Vals: 2}, Filter: gift.Gamma(2)},
		}
		result := ToFilters(filters)
		expected := []gift.Filter{
			gift.Contrast(2),
			gift.Gamma(2),
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Testing with gift.Filter", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := ToFilters(gift.Contrast(2))
		expected := []gift.Filter{gift.Contrast(2)}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Testing with unknown type", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		ToFilters(123)
	})
}
