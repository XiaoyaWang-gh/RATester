package internal

import (
	"context"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestAddTemplateFuncsNamespace_1(t *testing.T) {
	testCases := []struct {
		name string
		ns   func(d *deps.Deps) *TemplateFuncsNamespace
	}{
		{
			name: "Add a new namespace",
			ns: func(d *deps.Deps) *TemplateFuncsNamespace {
				return &TemplateFuncsNamespace{
					Name: "test",
					Context: func(ctx context.Context, v ...any) (any, error) {
						return "test", nil
					},
					OnCreated: func(namespaces map[string]any) {
						if _, ok := namespaces["test"]; !ok {
							t.Errorf("Expected namespace 'test' to be in the registry")
						}
					},
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			AddTemplateFuncsNamespace(tc.ns)
		})
	}
}
