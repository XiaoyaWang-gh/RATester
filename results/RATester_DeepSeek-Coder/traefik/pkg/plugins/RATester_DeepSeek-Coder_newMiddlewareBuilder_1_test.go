package plugins

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestNewMiddlewareBuilder_1(t *testing.T) {
	type args struct {
		ctx        context.Context
		goPath     string
		manifest   *Manifest
		moduleName string
		settings   Settings
	}
	tests := []struct {
		name    string
		args    args
		want    middlewareBuilder
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

			got, err := newMiddlewareBuilder(tt.args.ctx, tt.args.goPath, tt.args.manifest, tt.args.moduleName, tt.args.settings)
			if (err != nil) != tt.wantErr {
				t.Errorf("newMiddlewareBuilder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMiddlewareBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
