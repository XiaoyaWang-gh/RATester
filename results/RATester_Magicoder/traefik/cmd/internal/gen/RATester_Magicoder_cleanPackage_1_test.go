package main

import (
	"fmt"
	"path"
	"testing"
)

func TestCleanPackage_1(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want string
	}{
		{
			name: "github.com/traefik/paerser/types",
			src:  "github.com/traefik/paerser/types",
			want: "",
		},
		{
			name: "github.com/traefik/traefik/v3/pkg/tls",
			src:  "github.com/traefik/traefik/v3/pkg/tls",
			want: path.Join(destModuleName, destPkg, "tls"),
		},
		{
			name: "github.com/traefik/traefik/v3/pkg/types",
			src:  "github.com/traefik/traefik/v3/pkg/types",
			want: path.Join(destModuleName, destPkg, "types"),
		},
		{
			name: "other",
			src:  "other",
			want: "other",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := cleanPackage(tt.src); got != tt.want {
				t.Errorf("cleanPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}
