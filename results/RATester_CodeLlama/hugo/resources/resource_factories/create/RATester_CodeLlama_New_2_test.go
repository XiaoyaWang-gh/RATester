package create

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/resources"
)

func TestNew_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rs := &resources.Spec{}
	c := New(rs)

	assert.NotNil(t, c)
	assert.Equal(t, rs, c.rs)
	assert.NotNil(t, c.httpClient)
	assert.NotNil(t, c.httpCacheConfig)
	assert.NotNil(t, c.resourceIDDispatcher)
	assert.NotNil(t, c.remoteResourceChecker)
	assert.NotNil(t, c.remoteResourceLogger)
	assert.NotNil(t, c.cacheGetResource)
}
