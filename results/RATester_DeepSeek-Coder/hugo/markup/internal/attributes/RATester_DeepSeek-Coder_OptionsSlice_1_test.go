package attributes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOptionsSlice_1(t *testing.T) {
	tests := []struct {
		name string
		a    *AttributesHolder
		want []Attribute
	}{
		{
			name: "Test with empty options",
			a:    &AttributesHolder{},
			want: []Attribute{},
		},
		{
			name: "Test with non-empty options",
			a: &AttributesHolder{
				options: []Attribute{
					{Name: "option1", Value: "value1"},
					{Name: "option2", Value: "value2"},
				},
			},
			want: []Attribute{
				{Name: "option1", Value: "value1"},
				{Name: "option2", Value: "value2"},
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

			if got := tt.a.OptionsSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AttributesHolder.OptionsSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
