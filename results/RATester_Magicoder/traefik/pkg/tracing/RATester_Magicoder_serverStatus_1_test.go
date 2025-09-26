package tracing

import (
	"fmt"
	"reflect"
	"testing"

	"go.opentelemetry.io/otel/codes"
)

func TestServerStatus_1(t *testing.T) {
	tests := []struct {
		name  string
		code  int
		want  codes.Code
		want1 string
	}{
		{
			name:  "Test case 1",
			code:  200,
			want:  codes.Unset,
			want1: "",
		},
		{
			name:  "Test case 2",
			code:  500,
			want:  codes.Error,
			want1: "",
		},
		{
			name:  "Test case 3",
			code:  600,
			want:  codes.Error,
			want1: "Invalid HTTP status code 600",
		},
		{
			name:  "Test case 4",
			code:  99,
			want:  codes.Error,
			want1: "Invalid HTTP status code 99",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1 := serverStatus(tt.code)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serverStatus() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("serverStatus() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
