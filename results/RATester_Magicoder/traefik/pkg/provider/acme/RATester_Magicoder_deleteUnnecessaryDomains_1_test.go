package acme

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestDeleteUnnecessaryDomains_1(t *testing.T) {
	type args struct {
		ctx     context.Context
		domains []types.Domain
	}
	tests := []struct {
		name string
		args args
		want []types.Domain
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := deleteUnnecessaryDomains(tt.args.ctx, tt.args.domains); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteUnnecessaryDomains() = %v, want %v", got, tt.want)
			}
		})
	}
}
