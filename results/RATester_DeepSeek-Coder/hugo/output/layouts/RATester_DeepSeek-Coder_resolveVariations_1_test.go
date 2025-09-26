package layouts

import (
	"fmt"
	"reflect"
	"testing"
)

func TestResolveVariations_1(t *testing.T) {
	type fields struct {
		layoutVariations []string
		typeVariations   []string
		d                LayoutDescriptor
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
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

			l := &layoutBuilder{
				layoutVariations: tt.fields.layoutVariations,
				typeVariations:   tt.fields.typeVariations,
				d:                tt.fields.d,
			}
			if got := l.resolveVariations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("layoutBuilder.resolveVariations() = %v, want %v", got, tt.want)
			}
		})
	}
}
