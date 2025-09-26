package framework

import (
	"fmt"
	"testing"

	"github.com/bsm/gomega"
)

func TestExpectContainSubstring_1(t *testing.T) {
	gomega.RegisterTestingT(t)

	type args struct {
		actual  string
		substr  string
		explain []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				actual:  "Hello, world!",
				substr:  "world",
				explain: []interface{}{"Expecting the string to contain the substring"},
			},
		},
		{
			name: "Test case 2",
			args: args{
				actual:  "Hello, world!",
				substr:  "universe",
				explain: []interface{}{"Expecting the string to contain the substring"},
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

			ExpectContainSubstring(tt.args.actual, tt.args.substr, tt.args.explain...)
		})
	}
}
