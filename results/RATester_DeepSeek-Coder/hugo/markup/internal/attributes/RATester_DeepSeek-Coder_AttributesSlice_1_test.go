package attributes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAttributesSlice_1(t *testing.T) {
	tests := []struct {
		name string
		a    *AttributesHolder
		want []Attribute
	}{
		{
			name: "Test case 1",
			a: &AttributesHolder{
				attributes: []Attribute{
					{Name: "test1", Value: "value1"},
					{Name: "test2", Value: "value2"},
				},
			},
			want: []Attribute{
				{Name: "test1", Value: "value1"},
				{Name: "test2", Value: "value2"},
			},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.a.AttributesSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AttributesHolder.AttributesSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
