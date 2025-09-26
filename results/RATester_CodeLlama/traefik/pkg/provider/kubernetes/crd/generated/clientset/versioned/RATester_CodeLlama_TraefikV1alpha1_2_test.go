package versioned

import (
	"fmt"
	"testing"
)

func TestTraefikV1alpha1_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Clientset{}
	if got := c.TraefikV1alpha1(); got == nil {
		t.Error("TraefikV1alpha1() = nil")
	}
}
