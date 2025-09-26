package gateway

import (
	"fmt"
	"testing"

	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestBuildMatchRule_1(t *testing.T) {
	type args struct {
		hostnames []gatev1.Hostname
		match     gatev1.HTTPRouteMatch
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
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

			got, got1 := buildMatchRule(tt.args.hostnames, tt.args.match)
			if got != tt.want {
				t.Errorf("buildMatchRule() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("buildMatchRule() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
