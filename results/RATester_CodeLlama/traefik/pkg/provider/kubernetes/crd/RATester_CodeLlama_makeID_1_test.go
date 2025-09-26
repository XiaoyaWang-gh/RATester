package crd

import (
	"fmt"
	"testing"
)

func TestMakeID_1(t *testing.T) {
	var tests = []struct {
		namespace string
		name      string
		want      string
	}{
		{
			namespace: "test-namespace",
			name:      "test-name",
			want:      "test-namespace-test-name",
		},
		{
			namespace: "",
			name:      "test-name",
			want:      "test-name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := makeID(tt.namespace, tt.name); got != tt.want {
				t.Errorf("makeID() = %v, want %v", got, tt.want)
			}
		})
	}
}
