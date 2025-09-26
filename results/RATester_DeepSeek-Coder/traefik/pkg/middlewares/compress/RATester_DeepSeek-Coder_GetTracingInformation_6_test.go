package compress

import (
	"fmt"
	"reflect"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_6(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  string
		want2  trace.SpanKind
	}{
		{
			name: "Test case 1",
			fields: fields{
				name: "compress",
			},
			want:  "compress",
			want1: "compress",
			want2: trace.SpanKindInternal,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &compress{
				name: tt.fields.name,
			}
			got, got1, got2 := c.GetTracingInformation()
			if got != tt.want {
				t.Errorf("GetTracingInformation() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetTracingInformation() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GetTracingInformation() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
