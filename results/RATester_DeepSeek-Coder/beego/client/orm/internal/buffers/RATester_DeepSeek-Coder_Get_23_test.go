package buffers

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/valyala/bytebufferpool"
)

func TestGet_23(t *testing.T) {
	tests := []struct {
		name string
		want Buffer
	}{
		{
			name: "Test case 1",
			want: bytebufferpool.Get(),
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

			got := Get()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
