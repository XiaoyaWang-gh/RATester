package hugofs

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/loggers"
	"github.com/gohugoio/hugo/media"
	"github.com/spf13/afero"
)

func TestNewWalkway_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	var (
		fs  = afero.NewMemMapFs()
		cfg = WalkwayConfig{
			Fs: fs,
		}
	)

	w := NewWalkway(cfg)

	if w.cfg.Fs != fs {
		t.Errorf("cfg.Fs = %v, want %v", w.cfg.Fs, fs)
	}

	if w.cfg.PathParser != media.DefaultPathParser {
		t.Errorf("cfg.PathParser = %v, want %v", w.cfg.PathParser, media.DefaultPathParser)
	}

	if w.logger != loggers.NewDefault() {
		t.Errorf("w.logger = %v, want %v", w.logger, loggers.NewDefault())
	}
}
