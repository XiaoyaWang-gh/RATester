package migration

import (
	"fmt"
	"testing"
)

func TestAddOrUpdateRecord_1(t *testing.T) {
	type args struct {
		name   string
		status string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				name:   "test_migration",
				status: "down",
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				name:   "test_migration",
				status: "update",
			},
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				name:   "test_migration",
				status: "invalid_status",
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

			m := &Migration{}
			if err := m.addOrUpdateRecord(tt.args.name, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("Migration.addOrUpdateRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
