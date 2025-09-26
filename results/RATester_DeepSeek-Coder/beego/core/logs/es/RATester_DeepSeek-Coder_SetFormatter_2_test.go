package es

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/core/logs"
)

func TestSetFormatter_2(t *testing.T) {
	type args struct {
		f logs.LogFormatter
	}
	tests := []struct {
		name string
		el   *esLogger
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

			tt.el.SetFormatter(tt.args.f)
		})
	}
}
