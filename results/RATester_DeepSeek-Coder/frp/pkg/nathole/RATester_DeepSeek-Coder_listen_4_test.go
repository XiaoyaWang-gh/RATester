package nathole

import (
	"fmt"
	"testing"
)

func TestListen_4(t *testing.T) {
	type args struct {
		localAddr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test listen with valid local address",
			args: args{
				localAddr: "127.0.0.1:0",
			},
			wantErr: false,
		},
		{
			name: "Test listen with invalid local address",
			args: args{
				localAddr: "127.0.0.1:65536",
			},
			wantErr: true,
		},
		{
			name: "Test listen with empty local address",
			args: args{
				localAddr: "",
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

			_, err := listen(tt.args.localAddr)
			if (err != nil) != tt.wantErr {
				t.Errorf("listen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
