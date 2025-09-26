package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/output"
)

func TestisStandalone_1(t *testing.T) {
	tests := []struct {
		name string
		p    *pageMeta
		want bool
	}{
		{
			name: "Test case 1",
			p: &pageMeta{
				standaloneOutputFormat: output.Format{},
			},
			want: false,
		},
		{
			name: "Test case 2",
			p: &pageMeta{
				standaloneOutputFormat: output.Format{
					Name: "test",
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.p.isStandalone(); got != tt.want {
				t.Errorf("isStandalone() = %v, want %v", got, tt.want)
			}
		})
	}
}
