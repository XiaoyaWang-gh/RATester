package source

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/helpers"
	"github.com/spf13/afero"
)

func TestNewSourceSpec_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ps := &helpers.PathSpec{}
	fs := afero.NewMemMapFs()
	ss := NewSourceSpec(ps, nil, fs)
	assert.NotNil(t, ss)
}
