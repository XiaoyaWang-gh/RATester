package glob

import (
	"fmt"
	"testing"
)

func TestNewFilenameFilter_1(t *testing.T) {
	t.Run("should return error when invalid glob pattern", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := NewFilenameFilter([]string{"["}, []string{})
		if err == nil {
			t.Errorf("Expected error when invalid glob pattern")
		}
	})

	t.Run("should return nil and no error when no patterns provided", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		filter, err := NewFilenameFilter(nil, nil)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if filter != nil {
			t.Errorf("Expected nil filter when no patterns provided")
		}
	})

	t.Run("should create filter with inclusions", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		filter, err := NewFilenameFilter([]string{"*.go"}, []string{})
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if filter == nil {
			t.Errorf("Expected non-nil filter")
		}
		if len(filter.inclusions) != 1 {
			t.Errorf("Expected 1 inclusion, got %d", len(filter.inclusions))
		}
		if len(filter.dirInclusions) != 1 {
			t.Errorf("Expected 1 dir inclusion, got %d", len(filter.dirInclusions))
		}
	})

	t.Run("should create filter with exclusions", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		filter, err := NewFilenameFilter([]string{}, []string{"*.go"})
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if filter == nil {
			t.Errorf("Expected non-nil filter")
		}
		if len(filter.exclusions) != 1 {
			t.Errorf("Expected 1 exclusion, got %d", len(filter.exclusions))
		}
	})
}
