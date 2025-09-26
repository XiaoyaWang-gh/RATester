package env

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetAll_1(t *testing.T) {
	tests := []struct {
		name string
		want map[string]string
	}{
		{
			name: "Test case 1",
			want: map[string]string{
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

			if got := GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
