package framework

import (
	"fmt"
	"testing"

	"github.com/bsm/gomega"
)

func TestExpectEqualValues_1(t *testing.T) {
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
				actual:  1,
				extra:   1,
				explain: []interface{}{"Testing ExpectEqualValues"},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				actual:  "test",
				extra:   "test",
				explain: []interface{}{"Testing ExpectEqualValues"},
			},
		},
		{
			name: "Test Case 3",
			args: args{
				actual:  []int{1, 2, 3},
				extra:   []int{1, 2, 3},
				explain: []interface{}{"Testing ExpectEqualValues"},
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

			ExpectEqualValues(tt.args.actual, tt.args.extra, tt.args.explain...)
		})
	}
}
