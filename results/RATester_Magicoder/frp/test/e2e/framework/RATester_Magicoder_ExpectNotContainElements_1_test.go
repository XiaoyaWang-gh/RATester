package framework

import (
	"fmt"
	"testing"

	"github.com/bsm/gomega"
)

func TestExpectNotContainElements_1(t *testing.T) {
	gomega.RegisterTestingT(t)

	type args struct {
		actual  interface{}
		extra   interface{}
		explain []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				actual:  []int{1, 2, 3},
				extra:   []int{2, 3},
				explain: []interface{}{"Expect not to contain elements"},
			},
		},
		{
			name: "Test case 2",
			args: args{
				actual:  []int{1, 2, 3},
				extra:   []int{4, 5},
				explain: []interface{}{"Expect to contain elements"},
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

			ExpectNotContainElements(tt.args.actual, tt.args.extra, tt.args.explain...)
		})
	}
}
