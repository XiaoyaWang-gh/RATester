package plugins

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/yaegi/interp"
)

func TestNewYaegiMiddlewareBuilder_1(t *testing.T) {
	type args struct {
		i       *interp.Interpreter
		basePkg string
		imp     string
	}
	tests := []struct {
		name    string
		args    args
		want    *yaegiMiddlewareBuilder
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

			got, err := newYaegiMiddlewareBuilder(tt.args.i, tt.args.basePkg, tt.args.imp)
			if (err != nil) != tt.wantErr {
				t.Errorf("newYaegiMiddlewareBuilder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newYaegiMiddlewareBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
