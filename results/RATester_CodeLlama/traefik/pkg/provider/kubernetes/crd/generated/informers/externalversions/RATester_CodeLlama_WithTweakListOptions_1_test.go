package externalversions

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestWithTweakListOptions_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	tweakListOptions := func(options *v1.ListOptions) {
		options.LabelSelector = "app=traefik"
	}
	// when
	result := WithTweakListOptions(tweakListOptions)
	// then
	assert.NotNil(t, result)
}
