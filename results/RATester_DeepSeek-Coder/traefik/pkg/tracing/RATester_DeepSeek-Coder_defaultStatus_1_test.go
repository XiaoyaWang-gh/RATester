package tracing

import (
	"fmt"
	"reflect"
	"testing"

	"go.opentelemetry.io/otel/codes"
)

func TestDefaultStatus_1(t *testing.T) {
	tests := []struct {
		name  string
		code  int
		want  codes.Code
		want1 string
	}{
		{
			name:  "Test with valid HTTP status code",
			code:  200,
			want:  codes.Unset,
			want1: "",
		},
		{
			name:  "Test with invalid HTTP status code",
			code:  600,
			want:  codes.Error,
			want1: "Invalid HTTP status code 600",
		},
		{
			name:  "Test with HTTP status code 500",
			code:  500,
			want:  codes.Error,
			want1: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1 := defaultStatus(tt.code)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultStatus() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("defaultStatus() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
