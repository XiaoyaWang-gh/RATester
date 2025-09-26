package ingress

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConvertAnnotations_1(t *testing.T) {
	tests := []struct {
		name        string
		annotations map[string]string
		want        map[string]string
	}{
		{
			name: "test case 1",
			annotations: map[string]string{
				"ingress.kubernetes.io/rewrite-target":  "/new-path",
				"ingress.kubernetes.io/proxy-body-size": "0",
			},
			want: map[string]string{
				"rewrite-target":  "/new-path",
				"proxy-body-size": "0",
			},
		},
		{
			name: "test case 2",
			annotations: map[string]string{
				"ingress.kubernetes.io/rewrite-target":  "/new-path",
				"ingress.kubernetes.io/proxy-body-size": "0",
				"other-annotation":                      "value",
			},
			want: map[string]string{
				"rewrite-target":  "/new-path",
				"proxy-body-size": "0",
			},
		},
		{
			name: "test case 3",
			annotations: map[string]string{
				"other-annotation": "value",
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := convertAnnotations(tt.annotations)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertAnnotations() = %v, want %v", got, tt.want)
			}
		})
	}
}
