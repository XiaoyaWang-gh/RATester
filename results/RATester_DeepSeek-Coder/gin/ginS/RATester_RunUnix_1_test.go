package ginS

import (
	"fmt"
	"testing"
)

func TestRunUnix_1(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				file: "test.txt",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				file: "test2.txt",
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

			if err := RunUnix(tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("RunUnix() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
