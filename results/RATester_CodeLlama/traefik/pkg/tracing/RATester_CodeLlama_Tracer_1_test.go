package tracing

import (
	"fmt"
	"reflect"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestTracer_1(t *testing.T) {
	type args struct {
		name    string
		options []trace.TracerOption
	}
	tests := []struct {
		name string
		t    TracerProvider
		args args
		want trace.Tracer
	}{
		{
			name: "Returns the Traefik Tracer when the name is github.com/traefik/traefik",
			t: TracerProvider{
				tracer: &Tracer{},
			},
			args: args{
				name:    "github.com/traefik/traefik",
				options: nil,
			},
			want: &Tracer{},
		},
		{
			name: "Returns the wrapped TracerProvider when the name is not github.com/traefik/traefik",
			t: TracerProvider{
				tracer: &Tracer{},
			},
			args: args{
				name:    "github.com/traefik/traefik/v3",
				options: nil,
			},
			want: &Tracer{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.t.Tracer(tt.args.name, tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TracerProvider.Tracer() = %v, want %v", got, tt.want)
			}
		})
	}
}
