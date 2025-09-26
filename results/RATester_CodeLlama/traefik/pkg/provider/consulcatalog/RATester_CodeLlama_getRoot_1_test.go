package consulcatalog

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestGetRoot_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &connectCert{
		root: []string{"root1", "root2"},
	}
	result := c.getRoot()
	assert.Equal(t, []types.FileOrContent{"root1", "root2"}, result)
}
