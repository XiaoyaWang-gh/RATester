package migration

import (
	"fmt"
	"testing"
)

func TestMigrate_1(t *testing.T) {
	type args struct {
		migrationType string
	}
	tests := []struct {
		name string
		m    *Migration
		args args
	}{
		{
			name: "Test Migrate",
			m:    &Migration{},
			args: args{
				migrationType: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.m.Migrate(tt.args.migrationType)
			if tt.m.ModifyType != tt.args.migrationType {
				t.Errorf("Migrate() = %v, want %v", tt.m.ModifyType, tt.args.migrationType)
			}
		})
	}
}
