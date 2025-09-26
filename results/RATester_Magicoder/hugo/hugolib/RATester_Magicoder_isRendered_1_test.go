package hugolib

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestisRendered_1(t *testing.T) {
	tests := []struct {
		name string
		po   *pageOutput
		want bool
	}{
		{
			name: "Test case 1",
			po: &pageOutput{
				renderState: 1,
			},
			want: true,
		},
		{
			name: "Test case 2",
			po: &pageOutput{
				pco: &pageContentOutput{
					contentRendered: atomic.Bool{},
				},
			},
			want: true,
		},
		{
			name: "Test case 3",
			po: &pageOutput{
				renderState: 0,
				pco: &pageContentOutput{
					contentRendered: atomic.Bool{},
				},
			},
			want: false,
		},
		{
			name: "Test case 4",
			po: &pageOutput{
				renderState: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.po.isRendered(); got != tt.want {
				t.Errorf("isRendered() = %v, want %v", got, tt.want)
			}
		})
	}
}
