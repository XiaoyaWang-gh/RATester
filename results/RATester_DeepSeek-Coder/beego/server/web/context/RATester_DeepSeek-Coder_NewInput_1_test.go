package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewInput_1(t *testing.T) {
	tests := []struct {
		name string
		want *BeegoInput
	}{
		{
			name: "Test case 1",
			want: &BeegoInput{
				pnames:  make([]string, 0, maxParam),
				pvalues: make([]string, 0, maxParam),
				data:    make(map[interface{}]interface{}),
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

			got := NewInput()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
