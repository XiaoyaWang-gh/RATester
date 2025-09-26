package ingress

import (
	"fmt"
	"testing"

	netv1 "k8s.io/api/networking/v1"
)

func TestShouldProcessIngress_1(t *testing.T) {
	type args struct {
		ingress        *netv1.Ingress
		ingressClasses []*netv1.IngressClass
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

			p := &Provider{}
			if got := p.shouldProcessIngress(tt.args.ingress, tt.args.ingressClasses); got != tt.want {
				t.Errorf("shouldProcessIngress() = %v, want %v", got, tt.want)
			}
		})
	}
}
