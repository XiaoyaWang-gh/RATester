package testhelpers

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestWithServiceName_1(t *testing.T) {
	type args struct {
		serviceName string
	}
	tests := []struct {
		name string
		args args
		want func(*dynamic.Router)
	}{
		{
			name: "Test WithServiceName",
			args: args{
				serviceName: "testService",
			},
			want: func(r *dynamic.Router) {
				r.Service = "testService"
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := WithServiceName(tt.args.serviceName)
			router := &dynamic.Router{}
			got(router)
			if router.Service != tt.args.serviceName {
				t.Errorf("WithServiceName() = %v, want %v", router.Service, tt.args.serviceName)
			}
		})
	}
}
