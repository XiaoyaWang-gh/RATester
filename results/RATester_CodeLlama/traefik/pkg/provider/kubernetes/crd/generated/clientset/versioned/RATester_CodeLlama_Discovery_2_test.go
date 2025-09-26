package versioned

import (
	"fmt"
	"testing"
)

func TestDiscovery_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Clientset{}
	if got := c.Discovery(); got != nil {
		t.Errorf("Discovery() = %v, want nil", got)
	}
}
