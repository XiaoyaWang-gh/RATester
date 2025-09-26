package main

import (
	"fmt"
	"go/types"
	"reflect"
	"testing"
)

func TestNewCentrifuge_1(t *testing.T) {
	type args struct {
		rootPkg string
	}
	tests := []struct {
		name    string
		args    args
		want    *Centrifuge
		wantErr bool
	}{
		{
			name: "Test with valid package",
			args: args{
				rootPkg: "github.com/traefik/traefik/v3/pkg/config/dynamic",
			},
			want: &Centrifuge{
				TypeCleaner: func(typ types.Type, _ string) string {
					return typ.String()
				},
				PackageCleaner: func(s string) string {
					return s
				},
			},
			wantErr: false,
		},
		{
			name: "Test with invalid package",
			args: args{
				rootPkg: "github.com/traefik/traefik/v3/pkg/config/dynamic/invalid",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := NewCentrifuge(tt.args.rootPkg)
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
