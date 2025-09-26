package plugins

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/yaegi/interp"
)

func TestNewInterpreter_1(t *testing.T) {
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
			name: "test_newInterpreter_0",
			args: args{
				ctx:            context.Background(),
				goPath:         "",
				manifestImport: "",
			},
			want:    nil,
			wantErr: true,
		},
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
				t.Errorf("newInterpreter() got = %v, want %v", got, tt.want)
			}
		})
	}
}
