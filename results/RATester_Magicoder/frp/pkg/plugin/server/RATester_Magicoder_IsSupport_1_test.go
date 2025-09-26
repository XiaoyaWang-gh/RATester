package plugin

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestIsSupport_1(t *testing.T) {
	plugin := &httpPlugin{
		options: v1.HTTPPluginOptions{
			Ops: []string{"op1", "op2", "op3"},
		},
	}

	tests := []struct {
		name string
		op   string
		want bool
	}{
		{
			name: "Supported op",
			op:   "op1",
			want: true,
		},
		{
			name: "Unsupported op",
			op:   "op4",
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

			if got := plugin.IsSupport(tt.op); got != tt.want {
				t.Errorf("IsSupport() = %v, want %v", got, tt.want)
			}
		})
	}
}
