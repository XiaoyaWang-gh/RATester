package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/core/config"
)

func TestAssignSingleConfig_1(t *testing.T) {
	type args struct {
		p  interface{}
		ac config.Configer
	}
	tests := []struct {
		name string
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

			assignSingleConfig(tt.args.p, tt.args.ac)
		})
	}
}
