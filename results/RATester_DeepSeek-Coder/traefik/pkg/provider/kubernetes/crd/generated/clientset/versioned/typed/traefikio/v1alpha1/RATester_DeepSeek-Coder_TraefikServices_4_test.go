package v1alpha1

import (
	"fmt"
	"testing"

	"k8s.io/client-go/rest/fake"
)

func TestTraefikServices_4(t *testing.T) {
	client := &TraefikV1alpha1Client{
		restClient: &fake.RESTClient{},
	}

	ns := "default"

	t.Run("should return a TraefikServiceInterface", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := client.TraefikServices(ns)
		if got == nil {
			t.Errorf("TraefikServices() = %v, want not nil", got)
		}
	})
}
