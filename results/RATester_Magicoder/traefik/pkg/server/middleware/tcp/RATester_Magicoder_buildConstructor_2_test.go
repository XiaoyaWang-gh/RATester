package tcpmiddleware

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/tcp"
)

func TestBuildConstructor_2(t *testing.T) {
	type args struct {
		ctx            context.Context
		middlewareName string
	}
	tests := []struct {
		name    string
		b       *Builder
		args    args
		want    tcp.Constructor
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

			got, err := tt.b.buildConstructor(tt.args.ctx, tt.args.middlewareName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Builder.buildConstructor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Builder.buildConstructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
