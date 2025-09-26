package tracing

import (
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestNewTracing_1(t *testing.T) {
	type args struct {
		conf *static.Tracing
	}
	tests := []struct {
		name    string
		args    args
		want    *Tracer
		want1   io.Closer
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

			got, got1, err := NewTracing(tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTracing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTracing() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewTracing() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
