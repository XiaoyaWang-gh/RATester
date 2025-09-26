package requestdecorator

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestCnameResolve_1(t *testing.T) {
	type args struct {
		ctx        context.Context
		host       string
		resolvPath string
	}
	tests := []struct {
		name    string
		args    args
		want    *cnameResolv
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

			got, err := cnameResolve(tt.args.ctx, tt.args.host, tt.args.resolvPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("cnameResolve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cnameResolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
