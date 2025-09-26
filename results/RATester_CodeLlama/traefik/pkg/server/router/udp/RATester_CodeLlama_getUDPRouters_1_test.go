package udp

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestGetUDPRouters_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	entryPoints := []string{"foo", "bar"}

	m := &Manager{}

	got := m.getUDPRouters(ctx, entryPoints)

	want := make(map[string]map[string]*runtime.UDPRouterInfo)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("getUDPRouters() = %v, want %v", got, want)
	}
}
