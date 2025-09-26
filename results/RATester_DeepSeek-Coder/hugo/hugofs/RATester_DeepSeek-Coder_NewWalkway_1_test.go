package hugofs

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/loggers"
	"github.com/gohugoio/hugo/media"
	"github.com/spf13/afero"
)

func TestNewWalkway_1(t *testing.T) {
	t.Run("should panic if fs is not set", func(t *testing.T) {

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

		NewWalkway(WalkwayConfig{})
	})

	t.Run("should set default path parser if not set", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		fs := afero.NewMemMapFs()
		cfg := WalkwayConfig{Fs: fs}
		walkway := NewWalkway(cfg)

		if walkway.cfg.PathParser != media.DefaultPathParser {
			t.Errorf("Expected default path parser, got %v", walkway.cfg.PathParser)
		}
	})

	t.Run("should set default logger if not set", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		fs := afero.NewMemMapFs()
		cfg := WalkwayConfig{Fs: fs}
		walkway := NewWalkway(cfg)

		if walkway.logger == nil {
			t.Errorf("Expected default logger, got nil")
		}
	})

	t.Run("should return a new Walkway with the given config", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		fs := afero.NewMemMapFs()
		logger := loggers.NewDefault()
		cfg := WalkwayConfig{Fs: fs, Logger: logger}
		walkway := NewWalkway(cfg)

		if walkway.cfg.Fs != fs {
			t.Errorf("Expected fs to be %v, got %v", fs, walkway.cfg.Fs)
		}

		if walkway.logger != logger {
			t.Errorf("Expected logger to be %v, got %v", logger, walkway.logger)
		}
	})
}
