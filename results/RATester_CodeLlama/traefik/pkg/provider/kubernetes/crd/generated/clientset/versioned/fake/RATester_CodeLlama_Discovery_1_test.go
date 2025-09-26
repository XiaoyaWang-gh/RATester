package fake

import (
	"fmt"
	"testing"
)

func TestDiscovery_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Clientset{}
	want := c.discovery
	got := c.Discovery()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
