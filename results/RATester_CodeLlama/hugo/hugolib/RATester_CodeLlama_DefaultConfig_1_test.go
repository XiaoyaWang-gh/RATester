package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config/allconfig"
	"github.com/spf13/afero"
)

func TestDefaultConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := afero.NewMemMapFs()
	all, err := allconfig.LoadConfig(allconfig.ConfigSourceDescriptor{Fs: fs, Environ: []string{"none"}})
	if err != nil {
		t.Fatal(err)
	}
	if all.Base == nil {
		t.Fatal("Base is nil")
	}
}
