package externalversions

import (
	"fmt"
	"testing"

	cache "k8s.io/client-go/tools/cache"
)

func TestWithTransform_1(t *testing.T) {
	t.Run("WithTransform", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var transform cache.TransformFunc
		WithTransform(transform)
	})
}
