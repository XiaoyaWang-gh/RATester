package consulcatalog

import (
	"fmt"
	"testing"
)

func TestIsReady_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &connectCert{}
	if c.isReady() {
		t.Errorf("isReady() = true, want false")
	}

	c.root = []string{"root"}
	if c.isReady() {
		t.Errorf("isReady() = true, want false")
	}

	c.leaf.cert = "cert"
	if c.isReady() {
		t.Errorf("isReady() = true, want false")
	}

	c.leaf.key = "key"
	if !c.isReady() {
		t.Errorf("isReady() = false, want true")
	}
}
