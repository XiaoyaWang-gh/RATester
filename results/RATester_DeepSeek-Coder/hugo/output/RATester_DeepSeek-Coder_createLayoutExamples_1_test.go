package output

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateLayoutExamples_1(t *testing.T) {
	tests := []struct {
		name string
		want any
	}{
		{
			name: "Test createLayoutExamples",
			want: createLayoutExamples(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := createLayoutExamples()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createLayoutExamples() = %v, want %v", got, tt.want)
			}
		})
	}
}
