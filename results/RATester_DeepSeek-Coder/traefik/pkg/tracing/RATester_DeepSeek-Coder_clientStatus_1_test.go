package tracing

import (
	"fmt"
	"reflect"
	"testing"

	"go.opentelemetry.io/otel/codes"
)

func TestClientStatus_1(t *testing.T) {
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
			code:  700,
			want:  codes.Error,
			want1: "Invalid HTTP status code 700",
		},
		{
			name:  "Test with HTTP status code 400",
			code:  400,
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

			got, got1 := clientStatus(tt.code)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clientStatus() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("clientStatus() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
