package api

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestkeepService_1(t *testing.T) {
	type args struct {
		name      string
		item      *runtime.ServiceInfo
		criterion *searchCriterion
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_case_1",
			args: args{
				name: "service1",
				item: &runtime.ServiceInfo{
					Status: "enabled",
				},
				criterion: &searchCriterion{
					Status: "enabled",
				},
			},
			want: true,
		},
		{
			name: "test_case_2",
			args: args{
				name: "service2",
				item: &runtime.ServiceInfo{
					Status: "disabled",
				},
				criterion: &searchCriterion{
					Status: "enabled",
				},
			},
			want: false,
		},
		{
			name: "test_case_3",
			args: args{
				name: "service3",
				item: &runtime.ServiceInfo{
					Status: "enabled",
				},
				criterion: nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := keepService(tt.args.name, tt.args.item, tt.args.criterion); got != tt.want {
				t.Errorf("keepService() = %v, want %v", got, tt.want)
			}
		})
	}
}
