package framework

import (
	"fmt"
	"testing"

	"github.com/bsm/gomega"
)

func TestExpectConsistOf_1(t *testing.T) {
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
			name: "Test Case 1",
			args: args{
				actual:  []string{"a", "b", "c"},
				extra:   []string{"a", "b", "c"},
				explain: []interface{}{"Expecting slice to consist of the same elements"},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				actual:  []string{"a", "b", "c"},
				extra:   []string{"a", "b"},
				explain: []interface{}{"Expecting slice to consist of the same elements"},
			},
		},
		{
			name: "Test Case 3",
			args: args{
				actual:  []string{"a", "b", "c"},
				extra:   []string{"a", "b", "c", "d"},
				explain: []interface{}{"Expecting slice to consist of the same elements"},
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

			ExpectConsistOf(tt.args.actual, tt.args.extra, tt.args.explain...)
		})
	}
}
