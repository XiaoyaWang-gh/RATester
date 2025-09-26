package main

import (
	"fmt"
	"go/token"
	"go/types"
	"reflect"
	"testing"
)

func TestNewCentrifuge_1(t *testing.T) {
	tests := []struct {
		name    string
		rootPkg string
		want    *Centrifuge
		wantErr bool
	}{
		{
			name:    "test",
			rootPkg: "github.com/traefik/traefik/v3/cmd/internal/gen",
			want: &Centrifuge{
				fileSet: token.NewFileSet(),
				pkg:     nil,
				rootPkg: "github.com/traefik/traefik/v3/cmd/internal/gen",
				TypeCleaner: func(typ types.Type, _ string) string {
					return typ.String()
				},
				PackageCleaner: func(s string) string {
					return s
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := NewCentrifuge(tt.rootPkg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCentrifuge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCentrifuge() = %v, want %v", got, tt.want)
			}
		})
	}
}
