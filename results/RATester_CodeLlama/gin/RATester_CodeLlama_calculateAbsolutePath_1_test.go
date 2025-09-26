package gin

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestCalculateAbsolutePath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	group := &RouterGroup{
		basePath: "/v1",
	}
	assert.Equal(t, "/v1/test", group.calculateAbsolutePath("/test"))
}
