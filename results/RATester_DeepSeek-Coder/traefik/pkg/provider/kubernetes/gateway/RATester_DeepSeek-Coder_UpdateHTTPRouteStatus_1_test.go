package gateway

import (
	"context"
	"fmt"
	"testing"

	ktypes "k8s.io/apimachinery/pkg/types"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestUpdateHTTPRouteStatus_1(t *testing.T) {
	type args struct {
		ctx    context.Context
		route  ktypes.NamespacedName
		status gatev1.HTTPRouteStatus
	}
	tests := []struct {
		name    string
		c       *clientWrapper
		args    args
		wantErr bool
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

			if err := tt.c.UpdateHTTPRouteStatus(tt.args.ctx, tt.args.route, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("clientWrapper.UpdateHTTPRouteStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
