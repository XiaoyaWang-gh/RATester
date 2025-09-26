package commands

import (
	"fmt"
	"testing"
)

func TestimportFromJekyll_1(t *testing.T) {
	type args struct {
		args []string
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

			c := &importCommand{}
			if err := c.importFromJekyll(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("importFromJekyll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
