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
			name: "TestGet",
			want: bytebufferpool.Get(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
