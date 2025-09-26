package allconfig

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/docshelper"
)

func Testinit_35(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	allDecoderSetups := map[string]struct {
		internalOrDeprecated bool
	}{
		"test": {internalOrDeprecated: true},
	}

	docsProvider := func() docshelper.DocProvider {
		cfg := config.New()
		for configRoot, v := range allDecoderSetups {
			if v.internalOrDeprecated {
				continue
			}
			cfg.Set(configRoot, make(maps.Params))
		}
		lang := maps.Params{
			"en": maps.Params{
				"menus":  maps.Params{},
				"params": maps.Params{},
			},
		}
		cfg.Set("languages", lang)
		cfg.SetDefaultMergeStrategy()

		configHelpers := map[string]any{
			"mergeStrategy": cfg.Get(""),
		}
		return docshelper.DocProvider{"config_helpers": configHelpers}
	}

	docshelper.AddDocProviderFunc(docsProvider)

	// Add your test cases here
}
