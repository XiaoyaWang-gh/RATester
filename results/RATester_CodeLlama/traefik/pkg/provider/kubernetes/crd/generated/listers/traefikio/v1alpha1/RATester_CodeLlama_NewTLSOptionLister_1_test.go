package v1alpha1

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"k8s.io/client-go/tools/cache"
)

func TestNewTLSOptionLister_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var indexer cache.Indexer

	// when
	lister := NewTLSOptionLister(indexer)

	// then
	assert.NotNil(t, lister)
}
