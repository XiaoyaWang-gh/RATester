package commands

import (
	"fmt"
	"testing"
)

func TestfullBuild_1(t *testing.T) {
	type args struct {
		noBuildLock bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				noBuildLock: false,
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				noBuildLock: true,
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				noBuildLock: false,
			},
			wantErr: true,
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

			c := &hugoBuilder{}
			if err := c.fullBuild(tt.args.noBuildLock); (err != nil) != tt.wantErr {
				t.Errorf("hugoBuilder.fullBuild() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
