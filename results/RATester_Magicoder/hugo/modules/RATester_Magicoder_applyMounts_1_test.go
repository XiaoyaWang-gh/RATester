package modules

import (
	"fmt"
	"testing"
)

func TestapplyMounts_1(t *testing.T) {
	type args struct {
		moduleImport Import
		mod          *moduleAdapter
	}
	tests := []struct {
		name    string
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

			c := &collector{}
			if err := c.applyMounts(tt.args.moduleImport, tt.args.mod); (err != nil) != tt.wantErr {
				t.Errorf("applyMounts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
