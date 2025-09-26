package tracing

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestExtractCarrierIntoContext_1(t *testing.T) {
	type args struct {
		ctx     context.Context
		headers http.Header
	}
	tests := []struct {
		name string
		args args
		want context.Context
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

			if got := ExtractCarrierIntoContext(tt.args.ctx, tt.args.headers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractCarrierIntoContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
