package compress

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestConfigDefault_16(t *testing.T) {
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

	t.Run("should return the provided config when it is valid", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cfg := configDefault(Config{
			Level: LevelBestCompression,
		})
		assert.Equal(t, Config{
			Level: LevelBestCompression,
		}, cfg)
	})

	t.Run("should return the default config when the provided config is invalid", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cfg := configDefault(Config{
			Level: Level(-10),
		})
		assert.Equal(t, ConfigDefault, cfg)
	})
}
