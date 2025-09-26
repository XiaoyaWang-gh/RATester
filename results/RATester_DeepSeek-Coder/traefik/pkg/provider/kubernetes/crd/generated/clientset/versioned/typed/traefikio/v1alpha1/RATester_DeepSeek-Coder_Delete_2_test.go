package v1alpha1

import (
	"context"
	"fmt"
	"testing"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestDelete_2(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
		opts v1.DeleteOptions
	}
	tests := []struct {
		name    string
		c       *ingressRouteTCPs
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

			if err := tt.c.Delete(tt.args.ctx, tt.args.name, tt.args.opts); (err != nil) != tt.wantErr {
				t.Errorf("ingressRouteTCPs.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
