package consulcatalog

import (
	"fmt"
	"testing"
)

func TestItemServersTransportKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	item := itemData{
		ID:         "1234567890",
		Node:       "node1",
		Datacenter: "dc1",
		Name:       "name1",
		Namespace:  "namespace1",
		Address:    "127.0.0.1",
		Port:       "8080",
		Status:     "up",
		Labels:     map[string]string{"label1": "value1"},
		Tags:       []string{"tag1", "tag2"},
		ExtraConf:  configuration{Enable: true, ConsulCatalog: specificConfiguration{Connect: true, Canary: true}},
	}
	expected := "tls-namespace1-dc1-name1"
	actual := itemServersTransportKey(item)
	if actual != expected {
		t.Errorf("itemServersTransportKey() = %v, want %v", actual, expected)
	}
}
