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
		{
			name: "test",
			args: args{
				i:       &interp.Interpreter{},
				basePkg: "test",
				imp:     "test",
			},
			want: &yaegiMiddlewareBuilder{
				fnNew:          reflect.Value{},
				fnCreateConfig: reflect.Value{},
			},
			wantErr: false,
		},
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
				t.Errorf("newYaegiMiddlewareBuilder() got = %v, want %v", got, tt.want)
			}
		})
	}
}
