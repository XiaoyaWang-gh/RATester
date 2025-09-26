package modules

import (
	"fmt"
	"testing"
)

func TestApplyMounts_1(t *testing.T) {
	type args struct {
		moduleImport Import
		mod          *moduleAdapter
	}
	tests := []struct {
		name    string
		c       *collector
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

			if err := tt.c.applyMounts(tt.args.moduleImport, tt.args.mod); (err != nil) != tt.wantErr {
				t.Errorf("collector.applyMounts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
