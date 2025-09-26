package metrics

import (
	"fmt"
	"testing"
)

func TestNewProxy_4(t *testing.T) {
	type args struct {
		arg1 string
		arg2 string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				arg1: "test1",
				arg2: "test2",
			},
		},
		{
			name: "Test case 2",
			args: args{
				arg1: "test3",
				arg2: "test4",
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

			noopServerMetrics{}.NewProxy(tt.args.arg1, tt.args.arg2)
		})
	}
}
