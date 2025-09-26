package consulcatalog

import (
	"fmt"
	"testing"
)

func TestEquals_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &connectCert{
		root: []string{"root1", "root2"},
		leaf: keyPair{
			cert: "cert",
			key:  "key",
		},
	}
	other := &connectCert{
		root: []string{"root1", "root2"},
		leaf: keyPair{
			cert: "cert",
			key:  "key",
		},
	}
	if !c.equals(other) {
		t.Errorf("expected %v to equal %v", c, other)
	}
}
