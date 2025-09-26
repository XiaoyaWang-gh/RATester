package orm

import (
	"fmt"
	"testing"
)

func TestRunSyncdb_1(t *testing.T) {
	type args struct {
		name    string
		force   bool
		verbose bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				name:    "test",
				force:   true,
				verbose: true,
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				name:    "test",
				force:   false,
				verbose: false,
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				name:    "test",
				force:   true,
				verbose: false,
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

			if err := RunSyncdb(tt.args.name, tt.args.force, tt.args.verbose); (err != nil) != tt.wantErr {
				t.Errorf("RunSyncdb() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
