package basic

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/test/e2e/framework"
)

func TestRunClientServerTest_2(t *testing.T) {
	type args struct {
		f          *framework.Framework
		configures *generalTestConfigures
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

			runClientServerTest(tt.args.f, tt.args.configures)
		})
	}
}
