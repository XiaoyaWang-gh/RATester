package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewOutput_1(t *testing.T) {
	tests := []struct {
		name string
		want *BeegoOutput
	}{
		{
			name: "Test case 1",
			want: &BeegoOutput{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := NewOutput()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}
