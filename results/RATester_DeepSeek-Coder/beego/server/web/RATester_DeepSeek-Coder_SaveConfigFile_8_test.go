package web

import (
	"fmt"
	"testing"
)

func TestSaveConfigFile_8(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		b       *beegoAppConfig
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

			if err := tt.b.SaveConfigFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("beegoAppConfig.SaveConfigFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
