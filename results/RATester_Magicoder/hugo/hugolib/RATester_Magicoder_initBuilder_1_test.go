package hugolib

import (
	"fmt"
	"testing"

	"github.com/bep/logg"
	"github.com/gohugoio/hugo/config"
)

func TestinitBuilder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	b := NewIntegrationTestBuilder(IntegrationTestConfig{
		T: t,
		TxtarString: `
-- txtar --
a.txt: Hello, World!
--
`,
		BaseCfg:             config.New(),
		Environ:             []string{"FOO=BAR"},
		Running:             true,
		Watching:            true,
		Verbose:             true,
		LogLevel:            logg.LevelTrace,
		NeedsOsFS:           true,
		RunGC:               true,
		PrintAndKeepTempDir: true,
		NeedsNpmInstall:     true,
		NFDFormOnDarwin:     true,
		WorkingDir:          "/tmp/test",
		BuildCfg:            BuildCfg{},
	})

	err := b.initBuilder()
	if err != nil {
		t.Fatal(err)
	}
}
