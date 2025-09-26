package v1alpha1

import (
	"context"
	"fmt"
	"testing"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestDeleteCollection_2(t *testing.T) {
	type args struct {
		ctx      context.Context
		opts     v1.DeleteOptions
		listOpts v1.ListOptions
	}
	tests := []struct {
		name    string
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

			c := &ingressRouteTCPs{}
			if err := c.DeleteCollection(tt.args.ctx, tt.args.opts, tt.args.listOpts); (err != nil) != tt.wantErr {
				t.Errorf("ingressRouteTCPs.DeleteCollection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
