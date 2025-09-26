package opentelemetry

import (
	"fmt"
	"io"
	"reflect"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestSetup_1(t *testing.T) {
	type args struct {
		serviceName      string
		sampleRate       float64
		globalAttributes map[string]string
	}
	tests := []struct {
		name    string
		c       *Config
		args    args
		want    trace.Tracer
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

			got, got1, err := tt.c.Setup(tt.args.serviceName, tt.args.sampleRate, tt.args.globalAttributes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.Setup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.Setup() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Config.Setup() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
