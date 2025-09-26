package config

import (
	"fmt"
	"testing"
)

func TestRenderWithTemplate_1(t *testing.T) {
	t.Run("should return error if template parse fails", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := RenderWithTemplate([]byte("{{ .Invalid }}"), &Values{})
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("should return error if template execution fails", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := RenderWithTemplate([]byte("{{ .Envs.Invalid }}"), &Values{})
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("should render template with values", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := "test"
		values := &Values{
			Envs: map[string]string{
				"TEST": expected,
			},
		}
		out, err := RenderWithTemplate([]byte("{{ .Envs.TEST }}"), values)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if string(out) != expected {
			t.Errorf("expected %q, got %q", expected, string(out))
		}
	})
}
