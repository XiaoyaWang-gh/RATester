package versioned

import (
	"fmt"
	"testing"

	rest "k8s.io/client-go/rest"
)

func TestNewForConfigOrDie_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &rest.Config{}
	cs := NewForConfigOrDie(c)
	if cs == nil {
		t.Errorf("expected non-nil client")
	}
}
