package consulcatalog

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetName_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := itemData{
		ID:         "1",
		Node:       "node",
		Datacenter: "dc",
		Name:       "name",
		Namespace:  "namespace",
		Address:    "address",
		Port:       "port",
		Status:     "status",
		Labels:     map[string]string{"label": "label"},
		Tags:       []string{"tag"},
		ExtraConf:  configuration{Enable: true, ConsulCatalog: specificConfiguration{Connect: true, Canary: true}},
	}

	assert.Equal(t, "name-1", getName(i))
}
