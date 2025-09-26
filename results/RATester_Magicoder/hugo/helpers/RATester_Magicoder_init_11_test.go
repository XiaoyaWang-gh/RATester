package helpers

import (
	"fmt"
	"sort"
	"testing"

	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/gohugoio/hugo/docshelper"
)

func Testinit_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	docsProvider := func() docshelper.DocProvider {
		var chromaLexers []any

		sort.Sort(lexers.GlobalLexerRegistry.Lexers)

		for _, l := range lexers.GlobalLexerRegistry.Lexers {

			config := l.Config()

			lexerEntry := struct {
				Name    string
				Aliases []string
			}{
				config.Name,
				config.Aliases,
			}

			chromaLexers = append(chromaLexers, lexerEntry)

		}

		return docshelper.DocProvider{"chroma": map[string]any{"lexers": chromaLexers}}
	}

	docshelper.AddDocProviderFunc(docsProvider)
}
