package docshelper

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetDocProvider_1(t *testing.T) {
	tests := []struct {
		name string
		want DocProvider
	}{
		{
			name: "Test case 1",
			want: map[string]any{
				"key1": "value1",
				"key2": "value2",
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

			got := GetDocProvider()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDocProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}
