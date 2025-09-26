package migration

import (
	"fmt"
	"testing"
)

func TestaddOrUpdateRecord_1(t *testing.T) {
	type args struct {
		name   string
		status string
	}
	tests := []struct {
		name    string
		m       *Migration
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := tt.m.addOrUpdateRecord(tt.args.name, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("Migration.addOrUpdateRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
