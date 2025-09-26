package externalversions

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestWithNamespace_1(t *testing.T) {
	t.Run("WithNamespace", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Run("should set the namespace", func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			factory := &sharedInformerFactory{}
			WithNamespace("test-namespace")(factory)
			assert.Equal(t, "test-namespace", factory.namespace)
		})
	})
}
