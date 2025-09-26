package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/core/config"
)

func TestAssignConfig_1(t *testing.T) {
	type args struct {
		ac config.Configer
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

			if err := assignConfig(tt.args.ac); (err != nil) != tt.wantErr {
				t.Errorf("assignConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
