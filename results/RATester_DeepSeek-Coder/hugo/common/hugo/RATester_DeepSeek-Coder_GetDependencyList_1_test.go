package hugo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetDependencyList_1(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "Test case 1",
			want: []string{
				"github.com/example/dep1 v1.0.0",
				"github.com/example/dep2 v2.0.0",
				"github.com/example/dep3 v3.0.0",
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

			if got := GetDependencyList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDependencyList() = %v, want %v", got, tt.want)
			}
		})
	}
}
