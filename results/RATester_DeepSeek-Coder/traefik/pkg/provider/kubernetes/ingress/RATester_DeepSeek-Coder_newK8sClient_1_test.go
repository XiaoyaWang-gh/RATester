package ingress

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestNewK8sClient_1(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		p       *Provider
		args    args
		want    *clientWrapper
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

			got, err := tt.p.newK8sClient(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("newK8sClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newK8sClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
