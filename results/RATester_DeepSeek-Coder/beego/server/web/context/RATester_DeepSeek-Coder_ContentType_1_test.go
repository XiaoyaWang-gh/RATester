package context

import (
	"fmt"
	"testing"
)

func TestContentType_1(t *testing.T) {
	output := &BeegoOutput{}

	t.Run("Test with .json extension", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		output.ContentType(".json")
		if output.Context.ResponseWriter.Header().Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type to be application/json, got %s", output.Context.ResponseWriter.Header().Get("Content-Type"))
		}
	})

	t.Run("Test with json extension", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		output.ContentType("json")
		if output.Context.ResponseWriter.Header().Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type to be application/json, got %s", output.Context.ResponseWriter.Header().Get("Content-Type"))
		}
	})

	t.Run("Test with .xml extension", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		output.ContentType(".xml")
		if output.Context.ResponseWriter.Header().Get("Content-Type") != "application/xml" {
			t.Errorf("Expected Content-Type to be application/xml, got %s", output.Context.ResponseWriter.Header().Get("Content-Type"))
		}
	})

	t.Run("Test with xml extension", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		output.ContentType("xml")
		if output.Context.ResponseWriter.Header().Get("Content-Type") != "application/xml" {
			t.Errorf("Expected Content-Type to be application/xml, got %s", output.Context.ResponseWriter.Header().Get("Content-Type"))
		}
	})
}
