package commands

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config"
	"github.com/spf13/pflag"
)

func TestSetValueFromFlag_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	var flags *pflag.FlagSet
	var cfg config.Provider
	var key string
	var targetKey string
	var force bool

	setValueFromFlag(flags, key, cfg, targetKey, force)

	// TODO: add test cases.
}
