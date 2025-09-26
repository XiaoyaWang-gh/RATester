package lazy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew_59(t *testing.T) {
	tests := []struct {
		name string
		want *Init
	}{
		{
			name: "TestNew",
			want: &Init{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
