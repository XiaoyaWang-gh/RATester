package nomad

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetName_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := item{
		ID:         "1",
		Name:       "test",
		Namespace:  "default",
		Node:       "1",
		Datacenter: "us-west",
		Address:    "127.0.0.1",
		Port:       8080,
		Tags:       []string{"tag1", "tag2"},
		ExtraConf: configuration{
			Enable: true,
			Canary: true,
		},
	}

	assert.Equal(t, "test-1", getName(i))
}
