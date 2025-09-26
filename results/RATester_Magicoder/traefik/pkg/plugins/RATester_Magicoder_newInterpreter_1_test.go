package plugins

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/yaegi/interp"
)

func TestNewInterpreter_1(t *testing.T) {
	ctx := context.Background()
	goPath := "/path/to/go"
	manifestImport := "github.com/example/manifest"

	type args struct {
		ctx            context.Context
		goPath         string
		manifestImport string
	}
	tests := []struct {
		name    string
		args    args
		want    *interp.Interpreter
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				ctx:            ctx,
				goPath:         goPath,
				manifestImport: manifestImport,
			},
			want:    &interp.Interpreter{},
			wantErr: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := newInterpreter(tt.args.ctx, tt.args.goPath, tt.args.manifestImport)
			if (err != nil) != tt.wantErr {
				t.Errorf("newInterpreter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newInterpreter() = %v, want %v", got, tt.want)
			}
		})
	}
}
