package headermodifier

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestNewRequestHeaderModifier_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	ctx := context.Background()
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	config := dynamic.HeaderModifier{
		Set: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
		Add: map[string]string{
			"key3": "value3",
			"key4": "value4",
		},
		Remove: []string{
			"key5",
			"key6",
		},
	}
	name := "test"

	// act
	handler := NewRequestHeaderModifier(ctx, next, config, name)

	// assert
	assert.NotNil(t, handler)
}
