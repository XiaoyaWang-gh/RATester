package allconfig

import (
	"fmt"
	"testing"
)

func TestLoadConfig_2(t *testing.T) {
	type args struct {
		configName string
	}
	tests := []struct {
		name    string
		l       configLoader
		args    args
		want    string
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

			got, err := tt.l.loadConfig(tt.args.configName)
			if (err != nil) != tt.wantErr {
				t.Errorf("configLoader.loadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("configLoader.loadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
