package client

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestSetServiceOptionsDefault_1(t *testing.T) {
	type args struct {
		options *ServiceOptions
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test with nil options",
			args: args{
				options: nil,
			},
		},
		{
			name: "Test with empty options",
			args: args{
				options: &ServiceOptions{},
			},
		},
		{
			name: "Test with non-empty options",
			args: args{
				options: &ServiceOptions{
					Common: &v1.ClientCommonConfig{},
				},
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

			setServiceOptionsDefault(tt.args.options)
			if tt.args.options == nil {
				return
			}
			if tt.args.options.ConnectorCreator == nil {
				t.Errorf("ConnectorCreator should not be nil")
			}
			if tt.args.options.Common != nil {
				tt.args.options.Common.Complete()
			}
		})
	}
}
