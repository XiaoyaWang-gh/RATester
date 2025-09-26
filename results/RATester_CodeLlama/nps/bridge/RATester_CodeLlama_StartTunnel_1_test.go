package bridge

import (
	"fmt"
	"testing"
)

func TestStartTunnel_1(t *testing.T) {
	type args struct {
		s *Bridge
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_start_tunnel",
			args: args{
				s: &Bridge{
					TunnelPort: 10000,
				},
			},
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

			if err := tt.args.s.StartTunnel(); (err != nil) != tt.wantErr {
				t.Errorf("StartTunnel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
