package api

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestKeepTCPRouter_1(t *testing.T) {
	type args struct {
		name      string
		item      *runtime.TCPRouterInfo
		criterion *searchCriterion
	}
	tests := []struct {
		name string
		args args
		want bool
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

			if got := keepTCPRouter(tt.args.name, tt.args.item, tt.args.criterion); got != tt.want {
				t.Errorf("keepTCPRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
