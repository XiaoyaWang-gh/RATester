package observability

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestNewMiddleware_2(t *testing.T) {
	type args struct {
		next     http.Handler
		name     string
		typeName string
		spanKind trace.SpanKind
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
		{
			name: "Test case 1",
			args: args{
				next:     http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
				name:     "test",
				typeName: "test",
				spanKind: trace.SpanKindServer,
			},
			want: &middlewareTracing{
				next:     http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
				name:     "test",
				typeName: "test",
				spanKind: trace.SpanKindServer,
			},
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

			if got := NewMiddleware(tt.args.next, tt.args.name, tt.args.typeName, tt.args.spanKind); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
