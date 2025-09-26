package opentelemetry

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewFilterChainBuilder_2(t *testing.T) {
	tests := []struct {
		name string
		want *FilterChainBuilder
	}{
		{
			name: "TestNewFilterChainBuilder",
			want: &FilterChainBuilder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewFilterChainBuilder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterChainBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
