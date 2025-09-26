package crd

import (
	"fmt"
	"testing"
)

func TestMakeServersTransportKey_1(t *testing.T) {
	type args struct {
		parentNamespace      string
		serversTransportName string
	}
	tests := []struct {
		name    string
		c       *configBuilder
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test case 1",
			c: &configBuilder{
				allowCrossNamespace: true,
			},
			args: args{
				parentNamespace:      "default",
				serversTransportName: "test",
			},
			want:    "default_test",
			wantErr: false,
		},
		{
			name: "Test case 2",
			c: &configBuilder{
				allowCrossNamespace: false,
			},
			args: args{
				parentNamespace:      "default",
				serversTransportName: "default_test@kubernetescrd",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test case 3",
			c: &configBuilder{
				allowCrossNamespace: true,
			},
			args: args{
				parentNamespace:      "default",
				serversTransportName: "",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tt.c.makeServersTransportKey(tt.args.parentNamespace, tt.args.serversTransportName)
			if (err != nil) != tt.wantErr {
				t.Errorf("configBuilder.makeServersTransportKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("configBuilder.makeServersTransportKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
