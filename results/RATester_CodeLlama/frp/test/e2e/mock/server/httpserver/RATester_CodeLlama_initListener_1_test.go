package httpserver

import (
	"fmt"
	"testing"
)

func TestInitListener_1(t *testing.T) {
	type args struct {
		s *Server
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				s: &Server{
					bindAddr: "127.0.0.1",
					bindPort: 8080,
				},
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				s: &Server{
					bindAddr: "127.0.0.1",
					bindPort: 8080,
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

			if err := tt.args.s.initListener(); (err != nil) != tt.wantErr {
				t.Errorf("initListener() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
