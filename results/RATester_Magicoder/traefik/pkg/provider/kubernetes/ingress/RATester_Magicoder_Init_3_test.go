package ingress

import (
	"fmt"
	"testing"
)

func TestInit_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &Provider{
		Endpoint: "https://kubernetes.default.svc",
	}

	err := provider.Init()
	if err != nil {
		t.Errorf("Init() returned an error: %v", err)
	}
}
