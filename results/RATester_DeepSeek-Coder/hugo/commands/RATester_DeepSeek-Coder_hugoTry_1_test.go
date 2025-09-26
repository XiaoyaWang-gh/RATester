package commands

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/hugolib"
)

func TestHugoTry_1(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string
		want *hugolib.HugoSites
	}

	tests := []testCase{
		{
			name: "Test Case 1",
			want: &hugolib.HugoSites{},
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

			c := &hugoBuilder{}
			if got := c.hugoTry(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hugoBuilder.hugoTry() = %v, want %v", got, tt.want)
			}
		})
	}
}
