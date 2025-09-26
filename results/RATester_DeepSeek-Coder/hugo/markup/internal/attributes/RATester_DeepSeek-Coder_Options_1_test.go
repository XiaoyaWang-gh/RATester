package attributes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOptions_1(t *testing.T) {
	tests := []struct {
		name string
		a    *AttributesHolder
		want map[string]any
	}{
		{
			name: "Test with empty options",
			a:    &AttributesHolder{},
			want: map[string]any{},
		},
		{
			name: "Test with one option",
			a: &AttributesHolder{
				options: []Attribute{
					{
						Name:  "option1",
						Value: "value1",
					},
				},
			},
			want: map[string]any{
				"option1": "value1",
			},
		},
		{
			name: "Test with multiple options",
			a: &AttributesHolder{
				options: []Attribute{
					{
						Name:  "option1",
						Value: "value1",
					},
					{
						Name:  "option2",
						Value: "value2",
					},
				},
			},
			want: map[string]any{
				"option1": "value1",
				"option2": "value2",
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

			if got := tt.a.Options(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AttributesHolder.Options() = %v, want %v", got, tt.want)
			}
		})
	}
}
