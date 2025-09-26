package gateway

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestRegisterFilterFuncs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &Provider{}

	group := "testGroup"
	kind := "testKind"
	builderFunc := func(name, namespace string) (string, *dynamic.Middleware, error) {
		return "testName", &dynamic.Middleware{}, nil
	}

	provider.RegisterFilterFuncs(group, kind, builderFunc)

	if provider.groupKindFilterFuncs == nil {
		t.Fatal("groupKindFilterFuncs is nil")
	}

	if provider.groupKindFilterFuncs[group] == nil {
		t.Fatalf("group %s not found in groupKindFilterFuncs", group)
	}

	if provider.groupKindFilterFuncs[group][kind] == nil {
		t.Fatalf("kind %s not found in group %s in groupKindFilterFuncs", kind, group)
	}

	name, middleware, err := provider.groupKindFilterFuncs[group][kind]("testName", "testNamespace")
	if err != nil {
		t.Fatalf("error calling builderFunc: %v", err)
	}

	if name != "testName" {
		t.Fatalf("expected name %s, got %s", "testName", name)
	}

	if middleware == nil {
		t.Fatal("middleware is nil")
	}
}
