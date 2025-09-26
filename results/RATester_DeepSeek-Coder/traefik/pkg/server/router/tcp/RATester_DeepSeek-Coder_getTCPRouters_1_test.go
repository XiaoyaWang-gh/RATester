package tcp

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestGetTCPRouters_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	manager := &Manager{
		conf: &runtime.Configuration{},
	}

	entryPoints := []string{"http", "https"}
	expected := manager.conf.GetTCPRoutersByEntryPoints(ctx, entryPoints)

	result := manager.getTCPRouters(ctx, entryPoints)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
