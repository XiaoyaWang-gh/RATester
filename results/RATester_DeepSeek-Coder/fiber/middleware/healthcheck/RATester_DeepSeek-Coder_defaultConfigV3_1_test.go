package healthcheck

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/zeebo/assert"
)

func TestDefaultConfigV3_1(t *testing.T) {
	t.Parallel()

	t.Run("should return default config when no input config is provided", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cfg := defaultConfigV3()
		assert.Equal(t, defaultProbe, cfg.Probe)
	})

	t.Run("should return input config when input config is provided", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		customProbe := func(fiber.Ctx) bool { return false }
		cfg := defaultConfigV3(Config{Probe: customProbe})
		assert.Equal(t, customProbe, cfg.Probe)
	})

	t.Run("should return default config when input config has no Probe", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cfg := defaultConfigV3(Config{})
		assert.Equal(t, defaultProbe, cfg.Probe)
	})
}
