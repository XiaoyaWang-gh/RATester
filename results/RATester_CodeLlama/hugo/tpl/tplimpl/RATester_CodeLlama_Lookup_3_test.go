package tplimpl

import (
	"fmt"
	"sync"
	"testing"

	"github.com/alecthomas/assert"
	texttemplate "github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate"
)

func TestLookup_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	name := "test"
	tpl := &texttemplate.Template{}
	wrapper := &textTemplateWrapperWithLock{
		RWMutex:  &sync.RWMutex{},
		Template: tpl,
	}
	// when
	result, ok := wrapper.Lookup(name)
	// then
	assert.True(t, ok)
	assert.Equal(t, tpl, result)
}
