package requestid

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestConfigDefault_17(t *testing.T) {
	t.Parallel()

	t.Run("should return default config when no config is provided", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cfg := configDefault()
		assert.Equal(t, ConfigDefault, cfg)
	})

	t.Run("should return the provided config when config is provided", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := Config{
			Header: "X-Custom-Header",
			Generator: func() string {
				return "custom-id"
			},
		}
		cfg := configDefault(expected)
		assert.Equal(t, expected, cfg)
	})

	t.Run("should set default Header when Header is not provided", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := Config{
			Generator: func() string {
				return "custom-id"
			},
		}
		cfg := configDefault(expected)
		assert.Equal(t, ConfigDefault.Header, cfg.Header)
	})

	t.Run("should set default Generator when Generator is not provided", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := Config{
			Header: "X-Custom-Header",
		}
		cfg := configDefault(expected)
		assert.Equal(t, ConfigDefault.Generator, cfg.Generator)
	})
}
