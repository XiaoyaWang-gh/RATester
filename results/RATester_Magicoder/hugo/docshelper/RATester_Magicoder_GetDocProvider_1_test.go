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
			want: DocProvider{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name: "Test case 2",
			want: DocProvider{
				"key3": "value3",
				"key4": "value4",
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

			if got := GetDocProvider(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDocProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}
