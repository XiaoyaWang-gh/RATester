package ginS

import (
	"fmt"
	"testing"
)

func TestRun_1(t *testing.T) {
	type args struct {
		addr []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				addr: []string{"localhost:8080"},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				addr: []string{"invalid_address"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := Run(tt.args.addr...); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
