package proxy

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestGetName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pxy := &BaseProxy{
		name: "test_name",
	}
	assert.Equal(t, "test_name", pxy.GetName())
}
