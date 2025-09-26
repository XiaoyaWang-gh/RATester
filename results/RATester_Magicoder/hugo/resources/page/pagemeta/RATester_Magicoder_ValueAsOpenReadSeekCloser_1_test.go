package pagemeta

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/hugio"
)

func TestValueAsOpenReadSeekCloser_1(t *testing.T) {
	tests := []struct {
		name string
		s    Source
		want hugio.OpenReadSeekCloser
	}{
		{
			name: "test case 1",
			s:    Source{Value: "test value"},
			want: hugio.NewOpenReadSeekCloser(hugio.NewReadSeekerNoOpCloserFromString("test value")),
		},
		{
			name: "test case 2",
			s:    Source{Value: "another test value"},
			want: hugio.NewOpenReadSeekCloser(hugio.NewReadSeekerNoOpCloserFromString("another test value")),
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

			if got := tt.s.ValueAsOpenReadSeekCloser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Source.ValueAsOpenReadSeekCloser() = %v, want %v", got, tt.want)
			}
		})
	}
}
