package ginS

import (
	"fmt"
	"testing"
)

func TestRunFd_1(t *testing.T) {
	type args struct {
		fd int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				fd: 1,
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				fd: 2,
			},
			wantErr: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := RunFd(tt.args.fd); (err != nil) != tt.wantErr {
				t.Errorf("RunFd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
