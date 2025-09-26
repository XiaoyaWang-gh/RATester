package port

import (
	"fmt"
	"testing"
)

func TestUnmarshalFromName_1(t *testing.T) {
	t.Run("Test unmarshalFromName with valid name", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		name := "name1"
		builder, err := unmarshalFromName(name)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if builder.name != "name1" {
			t.Errorf("Expected name to be 'name1', got %v", builder.name)
		}
	})

	t.Run("Test unmarshalFromName with valid name and port range", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		name := "name1:1024-2048"
		builder, err := unmarshalFromName(name)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if builder.name != "name1" {
			t.Errorf("Expected name to be 'name1', got %v", builder.name)
		}
		if builder.rangePortFrom != 1024 {
			t.Errorf("Expected rangePortFrom to be 1024, got %v", builder.rangePortFrom)
		}
		if builder.rangePortTo != 2048 {
			t.Errorf("Expected rangePortTo to be 2048, got %v", builder.rangePortTo)
		}
	})

	t.Run("Test unmarshalFromName with invalid name", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		name := "name1:port"
		_, err := unmarshalFromName(name)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("Test unmarshalFromName with invalid port range", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		name := "name1:port-2048"
		_, err := unmarshalFromName(name)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
