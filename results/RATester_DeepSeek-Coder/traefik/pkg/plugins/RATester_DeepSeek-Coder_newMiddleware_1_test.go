package plugins

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewMiddleware_1(t *testing.T) {
	type args struct {
		config         map[string]interface{}
		middlewareName string
	}
	tests := []struct {
		name    string
		b       yaegiMiddlewareBuilder
		args    args
		want    pluginMiddleware
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

			got, err := tt.b.newMiddleware(tt.args.config, tt.args.middlewareName)
			if (err != nil) != tt.wantErr {
				t.Errorf("newMiddleware() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
