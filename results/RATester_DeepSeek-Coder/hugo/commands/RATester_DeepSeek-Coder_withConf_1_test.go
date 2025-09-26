package commands

import (
	"fmt"
	"testing"
)

func TestWithConf_1(t *testing.T) {
	type args struct {
		fn func(conf *commonConfig)
	}

	tests := []struct {
		name string
		c    *hugoBuilder
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &hugoBuilder{
				r:      tt.c.r,
				confmu: tt.c.confmu,
				conf:   tt.c.conf,
				s:      tt.c.s,
			}
			c.withConf(tt.args.fn)
		})
	}
}
