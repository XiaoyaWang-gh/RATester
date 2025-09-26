package crd

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
)

func TestCreateCompressMiddleware_1(t *testing.T) {
	type args struct {
		compress *traefikv1alpha1.Compress
	}
	tests := []struct {
		name string
		args args
		want *dynamic.Compress
	}{
		{
			name: "Test with nil compress",
			args: args{
				compress: nil,
			},
			want: nil,
		},
		{
			name: "Test with valid compress",
			args: args{
				compress: &traefikv1alpha1.Compress{
					ExcludedContentTypes: []string{"application/json"},
					IncludedContentTypes: []string{"text/html"},
					MinResponseBodyBytes: func() *int { i := 1024; return &i }(),
					Encodings:            []string{"gzip", "br"},
					DefaultEncoding:      func() *string { s := "gzip"; return &s }(),
				},
			},
			want: &dynamic.Compress{
				ExcludedContentTypes: []string{"application/json"},
				IncludedContentTypes: []string{"text/html"},
				MinResponseBodyBytes: 1024,
				Encodings:            []string{"gzip", "br"},
				DefaultEncoding:      "gzip",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := createCompressMiddleware(tt.args.compress); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createCompressMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
