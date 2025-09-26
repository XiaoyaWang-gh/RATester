package provider

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestMerge_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()

	configurations := map[string]*dynamic.Configuration{
		"config1": {
			HTTP: &dynamic.HTTPConfiguration{
				Routers: map[string]*dynamic.Router{
					"router1": {
						Rule: "Host(`example.com`)",
					},
				},
			},
		},
		"config2": {
			HTTP: &dynamic.HTTPConfiguration{
				Routers: map[string]*dynamic.Router{
					"router1": {
						Rule: "Host(`example.org`)",
					},
				},
			},
		},
	}

	expected := &dynamic.Configuration{
		HTTP: &dynamic.HTTPConfiguration{
			Routers: map[string]*dynamic.Router{
				"router1": {
					Rule: "Host(`example.org`)",
				},
			},
		},
	}

	result := Merge(ctx, configurations)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Merge() = %v, want %v", result, expected)
	}
}
