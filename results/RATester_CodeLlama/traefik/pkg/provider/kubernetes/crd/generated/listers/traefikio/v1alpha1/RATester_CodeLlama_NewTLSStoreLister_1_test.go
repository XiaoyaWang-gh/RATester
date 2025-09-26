package v1alpha1

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"k8s.io/client-go/tools/cache"
)

func TestNewTLSStoreLister_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var indexer cache.Indexer

	// when
	tLSStoreLister := NewTLSStoreLister(indexer)

	// then
	assert.NotNil(t, tLSStoreLister)
}
