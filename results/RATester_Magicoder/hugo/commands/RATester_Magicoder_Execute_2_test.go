package commands

import (
	"fmt"
	"testing"
)

func TestExecute_2(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				args: []string{"arg1", "arg2"},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				args: []string{"arg3", "arg4"},
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

			if err := Execute(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
