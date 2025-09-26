package config

import (
	"fmt"
	"testing"
)

func TestSaveConfigFile_2(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		c       *IniConfigContainer
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

			c := &IniConfigContainer{}
			if err := c.SaveConfigFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("IniConfigContainer.SaveConfigFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
