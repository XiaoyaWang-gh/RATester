package v1alpha1

import (
	"fmt"
	"testing"

	scheme "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

func TestSetConfigDefaults_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := &rest.Config{}
	err := setConfigDefaults(config)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if config.GroupVersion != nil {
		t.Errorf("unexpected group version: %v", config.GroupVersion)
	}
	if config.APIPath != "/apis" {
		t.Errorf("unexpected api path: %v", config.APIPath)
	}
	if config.NegotiatedSerializer != scheme.Codecs.WithoutConversion() {
		t.Errorf("unexpected negotiated serializer: %v", config.NegotiatedSerializer)
	}
	if config.UserAgent != rest.DefaultKubernetesUserAgent() {
		t.Errorf("unexpected user agent: %v", config.UserAgent)
	}
}
