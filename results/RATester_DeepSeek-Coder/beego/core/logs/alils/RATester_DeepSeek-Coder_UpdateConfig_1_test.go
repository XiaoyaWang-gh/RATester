package alils

import (
	"fmt"
	"testing"
)

func TestUpdateConfig_1(t *testing.T) {
	type args struct {
		c *LogConfig
	}
	tests := []struct {
		name    string
		p       *LogProject
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

			if err := tt.p.UpdateConfig(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("LogProject.UpdateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
